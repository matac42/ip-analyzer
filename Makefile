setup:
	-@multipass stop dev-ip-analyzer
	-@multipass delete dev-ip-analyzer
	-@multipass purge
	multipass launch --name dev-ip-analyzer --mem 4G --disk 10G --cpus 2 --mount .:/home/ubuntu/ip-analyzer 22.04
	multipass exec dev-ip-analyzer sudo sh go-setup.sh
	multipass exec dev-ip-analyzer sudo sh namespaces/router.sh