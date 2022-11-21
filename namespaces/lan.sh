ip -all netns del

ip netns add ns1
ip netns add ns2
ip netns add ns3
ip link add v-net-0 type bridge
ip link set ns1 v-net-0 up

ip link add veth-ns1 type veth peer name veth-ns1-br
ip link add veth-ns2 type veth peer name veth-ns2-br
ip link add veth-ns3 type veth peer name veth-ns3-br
ip link set veth-ns1 netns ns1
ip link set veth-ns1-br master v-net-0
ip link set veth-ns2 netns ns2
ip link set veth-ns2-br master v-net-0
ip link set veth-ns3 netns ns3
ip link set veth-ns3-br master v-net-0
ip -n ns1 addr add 192.168.1.1/24 dev veth-ns1
ip -n ns2 addr add 192.168.1.2/24 dev veth-ns2
ip -n ns3 addr add 192.168.1.3/24 dev veth-ns3
ip -n ns1 link set veth-ns1 up
ip -n ns2 link set veth-ns2 up
ip -n ns3 link set veth-ns3 up
ip link set veth-ns1-br up
ip link set veth-ns2-br up
ip link set veth-ns3-br up
ip -n ns1 link set lo up
ip -n ns2 link set lo up
ip -n ns3 link set lo up

ip netns exec ns1 ping 192.168.1.2