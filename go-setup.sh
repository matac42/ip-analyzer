add-apt-repository -y ppa:longsleep/golang-backports
apt-get update
apt-get install -y golang iproute2 iputils-ping tcpdump telnet dnsutils traceroute netcat curl \
        ldap-utils postgresql-client freeradius vim arping nmap ncat net-tools bridge-utils