IMAGE := edwardi/scratch_demo
VERSION := 0.1
DOCKER_IMAGE := ${IMAGE}:${VERSION}

# DIR is an example of defining a target-specific variable
print: DIR := "Something else"
print:
	@echo Example target-variable ${DIR}
	@echo DOCKER_IMAGE=${DOCKER_IMAGE}

push:
	docker image push ${DOCKER_IMAGE}

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -o app .
	docker image build -t ${DOCKER_IMAGE} .

run-detached:
	docker run -d -h MyHostName ${DOCKER_IMAGE}
	@echo 'you can now call `make inspect` and see your container running'
	@echo 'can call `make run` to see it logging interactively

run:
	docker run -it -h dockerAppHost ${DOCKER_IMAGE}

inspect:
	@docker container inspect `docker ps -q --filter ancestor=${DOCKER_IMAGE}` | jq

deploy:
	kubectl apply -f deployment.yaml

destroy:
	kubectl delete deployment demo