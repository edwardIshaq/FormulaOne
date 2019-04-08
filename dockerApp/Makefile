IMAGE := docker-app
NETWORK := my-bridge
CONTAINER := goServer

build-image:
	@docker image build -t docker-app .

run:
	@docker run -it \
	-h dockerAppHost \
	docker-app

attach:
	@docker exec -it ${CONTAINER} bash