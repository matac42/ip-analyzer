# 一度全てのネットワーク名前空間を削除する
ip --all netns delete

# 名前空間を作成する
ip netns add node1
ip netns add node2
ip netns add node3
ip netns add bridge

# 仮想ネットワークインターフェースを作成する
ip link add node1-veth0 type veth peer name node1-br0
ip link add node2-veth0 type veth peer name node2-br0
ip link add node3-veth0 type veth peer name node3-br0

# インタフェースを名前空間に所属させる
ip link set node1-veth0 netns node1
ip link set node2-veth0 netns node2
ip link set node3-veth0 netns node3
ip link set node1-br0 netns bridge
ip link set node2-br0 netns bridge
ip link set node3-br0 netns bridge

# インタフェースを有効にする
ip netns exec node1 ip link set node1-veth0 up
ip netns exec node2 ip link set node2-veth0 up
ip netns exec node3 ip link set node3-veth0 up
ip netns exec bridge ip link set node1-br0 up
ip netns exec bridge ip link set node2-br0 up
ip netns exec bridge ip link set node3-br0 up
ip netns exec node1 ip link set lo up
ip netns exec node2 ip link set lo up
ip netns exec node3 ip link set lo up

# アドレスを設定する
ip netns exec node1 ip a add 192.168.1.1/24 dev node1-veth0 
ip netns exec node2 ip a add 192.168.1.2/24 dev node2-veth0 
ip netns exec node3 ip a add 192.168.1.3/24 dev node3-veth0 

# ブリッジ(ハブ)を作成する
ip netns exec bridge ip link add dev br0 type bridge
ip netns exec bridge ip link set br0 up

# インタフェースをブリッジに接続する
ip netns exec bridge ip link set node1-br0 master br0
ip netns exec bridge ip link set node2-br0 master br0
ip netns exec bridge ip link set node3-br0 master br0