IMAGE := scratch_demo

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -o app .
	docker image build -t ${IMAGE} .

run-detached:
	docker run -d -h MyHostName ${IMAGE}
	@echo 'you can now call `make inspect` and see your container running'
	@echo 'can call `make run` to see it logging interactively

run:
	docker run -it -h dockerAppHost ${IMAGE}

inspect:
	@docker container inspect `docker ps -q --filter ancestor=${IMAGE}` | jq