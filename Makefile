GOVERSION := 1.10
SOURCE=$(shell find . -name '*.go')

# build a binary for linux-amd64
build: $(SOURCE)
	docker run --rm -v $(shell pwd):/go/src/github.com/giantswarm/kibana-sidecar \
		-e GOPATH=/go -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=1 \
		-w /go/src/github.com/giantswarm/kibana-sidecar \
		golang:$(GOVERSION)-alpine go build -o ./kibana-sidecar

docker-build: build
	docker build -t quay.io/giantswarm/kibana-sidecar .
