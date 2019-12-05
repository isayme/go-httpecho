APP_NAME := httpecho
APP_VERSION := $(shell git describe --tags || git rev-parse HEAD)
APP_PKG := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

.PHONY: image
image:
	docker build \
	--build-arg APP_NAME=${APP_NAME} \
	--build-arg APP_VERSION=${APP_VERSION} \
	--build-arg APP_PKG=${APP_PKG} \
	-t ${APP_NAME}:${APP_VERSION} .

.PHONY: dev
dev:
	go run main.go
