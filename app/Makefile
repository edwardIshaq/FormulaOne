NETWORK := my-bridge
CONTAINER := go_server
IMAGE := edwrardi/go_server
VERSION := 1.2
IMAGE_FULL := ${IMAGE}:${VERSION}

print:
	@echo IMAGE_FULL=${IMAGE_FULL}
	@echo NETWORK=${NETWORK}
	@echo CONTAINER=${CONTAINER}

build:
	@docker image build -t ${IMAGE}:${VERSION} .

run-docker:
	docker run \
	--name ${CONTAINER} \
	--network ${NETWORK} \
	-d ${IMAGE_FULL}

attach:
	@docker exec -it ${CONTAINER} bash