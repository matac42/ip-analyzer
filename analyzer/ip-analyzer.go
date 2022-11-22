package analyzer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/netip"
	"regexp"
	"time"

	"github.com/mdlayher/arp"
)

type Result struct {
	ID  int
	IP  string
	MAC string
}

func (r *Result) SetResult(ip string, mac string) {
	r.IP = ip
	r.MAC = mac
}

func (r *Result) GetResult() (string, string) {
	return r.IP, r.MAC
}

var (
	duration = 1000 * time.Microsecond

	iface = "node2-veth0"

	networkIP = "192.168.1.0/24"
)

func Analyze() string {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Request hardware address for IP address
	re := regexp.MustCompile(`\/\d*`)
	networkAddress := re.ReplaceAllString(networkIP, "")
	fmt.Println("networkAddress: ", networkAddress)

	re = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\/`)
	prefix := re.ReplaceAllString(networkIP, "")
	fmt.Println("prefix: ", prefix)

	ip, err := netip.ParseAddr(networkAddress)
	if err != nil {
		log.Fatal(err)
	}

	broadcast, err := netip.ParseAddr("192.168.1.255")
	if err != nil {
		log.Fatal(err)
	}

	ip = ip.Next()

	for ip.Less(broadcast) {
		ifi, err := net.InterfaceByName(iface)
		if err != nil {
			log.Fatal(err)
		}

		c, err := arp.Dial(ifi)
		if err != nil {
			log.Fatal(err)
		}

		if err := c.SetDeadline(time.Now().Add(duration)); err != nil {
			log.Fatal(err)
		}

		mac, err := c.Resolve(ip)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			result, err := db.Exec("INSERT INTO ip_analyzer (ip, mac) VALUES ($1, $2) RETURNING id", ip, mac)
			if err != nil {
				log.Fatal(err)
			}
			log.Default().Println(result)
		}
		err = c.Close()
		if err != nil {
			log.Fatal(err)
		}
		ip = ip.Next()
	}

	return ""
}

func ArpResult() string {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT ip, mac FROM ip_analyzer")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var ip string
	var mac string

	var result []Result

	for rows.Next() {
		switch err := rows.Scan(&id, &ip, &mac); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Printf("id: %d, ip: %s, mac: %s", id, ip, mac)

			result = append(result, Result{ID: id, IP: ip, MAC: mac})
		default:
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	resultJSON, err := json.Marshal(result)

	return string(resultJSON)
}

func GetProgress() string {
	return ""
}
