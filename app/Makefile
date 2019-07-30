NETWORK := my-bridge
CONTAINER := goServer
IMAGE := edwrardi/goServer
VERSION := 1.2
IMAGE_FULL := ${IMAGE}:${VERSION}

print: 
	@echo IMAGE_FULL=${IMAGE_FULL}
	@echo NETWORK=${NETWORK}
	@echo CONTAINER=${CONTAINER}

build:
	@docker image build -t ${IMAGE}:${VERSION} .

run:
	docker run \
	--name ${CONTAINER} \
	--network ${NETWORK} \
	-d ${IMAGE}:${VERSION}

attach:
	@docker exec -it ${CONTAINER} bash