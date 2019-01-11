VERSION=snapshot
REVISION=$(shell git rev-parse --short HEAD)

GOXCMD=gox
DEPCMD=dep

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=rancher-deployer
ARCH=amd64
OS=darwin linux
LDFLAGS=-s -w -X main.version=$(VERSION) -X main.revision=$(REVISION)

all: clean deps build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf dist

run:
	CGO_ENABLED=0 $(GOBUILD) -ldflags="${LDFLAGS}" -o $(BINARY_NAME) -v
	./$(BINARY_NAME) $(ARGS)

deps:
	command -v dep >/dev/null 2>&1 || go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/mitchellh/gox
	$(DEPCMD) ensure

release: clean
	CGO_ENABLED=0 $(GOXCMD) -arch="${ARCH}" -os="${OS}" -ldflags="${LDFLAGS}" -output="dist/${BINARY_NAME}_{{.OS}}_x86_64"