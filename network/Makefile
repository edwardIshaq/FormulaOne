# Network details
NETWORK_NAME := "my-bridge"
NETWORK_DRIVER := "bridge"
NETWORK_SUBNET := "11.0.0.1/24"

# you can override using CLI arguments :
# make NETWORK_NAME="MyNetworkName" print

print:
	@echo NETWORK_NAME=${NETWORK_NAME}
	@echo NETWORK_DRIVER=${NETWORK_DRIVER}
	@echo NETWORK_SUBNET=${NETWORK_SUBNET}

build:
	docker network create -d ${NETWORK_DRIVER} --subnet ${NETWORK_SUBNET} ${NETWORK_NAME}

inspect:
	docker network inspect ${NETWORK_NAME}