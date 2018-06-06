GOXCMD=gox
DEPCMD=dep

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=rancher-deployer
ARCH=amd64 386
OS=darwin linux
LDFLAGS=-s -w

all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf dist

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(DEPCMD) ensure

release: clean
	$(GOXCMD) -arch="${ARCH}" -os="${OS}" -ldflags="${LDFLAGS}" -output="dist/${BINARY_NAME}_{{.OS}}_{{.Arch}}"