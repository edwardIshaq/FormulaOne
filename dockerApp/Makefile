IMAGE := scratch_demo
NETWORK := my-bridge
CONTAINER := goServer

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -o app .
	docker image build -t ${IMAGE} .

run:
	@docker run -it \
	-h dockerAppHost \
	${IMAGE}

attach:
	@docker exec -it ${CONTAINER} bash