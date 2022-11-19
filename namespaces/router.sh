ip -all netns del

ip netns add ns1
ip netns add router
ip netns add ns2

ip link add ns1-veth0 type veth peer name gw-veth0
ip link add ns2-veth0 type veth peer name gw-veth1

ip link set ns1-veth0 netns ns1
ip link set gw-veth0 netns router
ip link set gw-veth1 netns router
ip link set ns2-veth0 netns ns2

ip netns exec ns1 ip link set ns1-veth0 up
ip netns exec router ip link set gw-veth0 up
ip netns exec router ip link set gw-veth1 up
ip netns exec ns2 ip link set ns2-veth0 up

ip netns exec ns1 ip address add 192.168.2.1/24 dev ns1-veth0
ip netns exec router ip address add 192.168.2.254/24 dev gw-veth0
ip netns exec router ip address add 192.168.3.254/24 dev gw-veth1
ip netns exec ns2 ip address add 192.168.3.1/24 dev ns2-veth0

ip netns exec ns1 ip route add default via 192.168.2.254
ip netns exec ns2 ip route add default via 192.168.3.254

ip netns exec router sysctl net.ipv4.ip_forward=1
ip netns exec router sysctl -w net.ipv4.ping_group_range='0 2147483647'
ip netns exec ns1 sysctl -w net.ipv4.ping_group_range='0 2147483647'
ip netns exec ns2 sysctl -w net.ipv4.ping_group_range='0 2147483647'

ip netns exec ns1 ping -c 2 192.168.2.254 -I 192.168.2.1
ip netns exec ns2 ping -c 2 192.168.3.254 -I 192.168.3.1
ip netns exec ns1 ping -c 2 192.168.3.1 -I 192.168.2.1