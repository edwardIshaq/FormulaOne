# For a shorter version
# git rev-parse --short HEAD
IMAGE_TAG := "edwardi/go_server"
VERSION := $(shell git rev-parse --short HEAD)

BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# $ date -R
# Wed, 11 Jul 2018 10:25:26 +0800

# $ date -u +"%Y-%m-%dT%H:%M:%SZ"
# 2018-07-11T02:25:23Z

VCS_URL := $(shell git config --get remote.origin.url)

# $(shell git log -1 --pretty=%h)
VCS_REF := $(shell git rev-parse HEAD)

NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)

print:
	@echo IMAGE_TAG=${IMAGE_TAG}
	@echo VERSION=${VERSION}
	@echo BUILD_DATE=${BUILD_DATE}
	@echo VCS_URL=${VCS_URL}
	@echo VCS_REF=${VCS_REF}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}

build:
	docker build \
	--build-arg VERSION="${VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg VCS_URL="${VCS_URL}" \
	--build-arg VCS_REF="${VCS_REF}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" \
	-t ${IMAGE_TAG} .

run:
	@docker run -it -p 8080:8080 --network my-bridge ${IMAGE_TAG}

label:
	@docker image inspect --format='{{json .Config.Labels}}' ${IMAGE_TAG}

test:
	@echo ${IMAGE_TAG}