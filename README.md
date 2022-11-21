# IP Analyzer

This project is incomplete.

- [x] Detects IP addresses used on the LAN.
- [ ] Identify the intended use of the host that has the IP.
- [ ] Notify duplicate IPs.

# For Developers

I am developing on an Ubuntu 22.04 VM built with Multipass.
Here is how to set up a development environment using [Multipass](https://multipass.run/).

```
cd ip-analyzer
multipass launch --name dev-ip-analyzer --mem 2G --disk 10G --cpus 2 --mount .:/home/ubuntu/ip-analyzer 22.04
multipass exec dev-ip-analyzer sudo sh go-setup.sh
multipass exec dev-ip-analyzer sudo sh namespaces/router.sh
```

If you build the `./ping` command, you will see it work like this command.

```
multipass exec dev-ip-analyzer -- sudo ip netns exec ns1 ./ping 192.168.3.1 -I 192.168.2.1
```

If you need to create a new namespace, create a new script in `./namespaces`.

When developing with a VM other than Multipass, build the environment as follows.

```
cd ip-analyzer
sudo sh go-setup.sh
sudo sh namespaces/router.sh
```
