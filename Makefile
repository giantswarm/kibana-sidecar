GOVERSION := 1.10
SOURCE=$(shell find . -name '*.go')

docker-build:
	docker build -t quay.io/giantswarm/kibana-sidecar .
