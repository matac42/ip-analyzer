package analyzer

import (
	"fmt"
	"log"
	"net"
	"net/netip"
	"regexp"
	"time"

	"github.com/mdlayher/arp"
)

var (
	duration = 1000 * time.Microsecond

	iface = "node2-veth0"

	networkIP = "192.168.1.0/24"
)

func main() {

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
			// fmt.Println("error: ", err)
		} else {
			fmt.Printf("%s -> %s\n", ip, mac)
		}
		err = c.Close()
		if err != nil {
			log.Fatal(err)
		}
		ip = ip.Next()
	}
}
