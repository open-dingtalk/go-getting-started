ifndef GOCMD
    GOCMD=go
endif

ifndef GOBUILD
    GOBUILD=$(GOCMD) build
endif

GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
BINARY_NAME=coolapp-go-getting-start

all: build

build:
	$(GOBUILD) -o _build/$(BINARY_NAME) -v ./cmd
test:
	mkdir cover && $(GOTEST) -v -cover=true -coverprofile=cover/go-getting-start.cover ./...
	go tool cover -html=cover/go-getting-start.cover -o cover/go-getting-start.html
clean:
	$(GOCLEAN)
	rm -f _build/*
run:
	$(GOBUILD) -o _build/$(BINARY_NAME) -v
	_build/$(BINARY_NAME)
fmt:
	find ./ -name "*.go" | xargs gofmt -w
