BINARY = bin/$(WORKINGDIR)
BUILD_FLAGS = -ldflags="-s -w" 
CMD=main.go
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
IMAGE= codebaker/$(WORKINGDIR):$(TAG)
TAG=$(shell git describe)
UNAME= $(shell uname | tr '[:upper:]' '[:lower:]')
WORKINGDIR= $(shell pwd | rev | cut -d'/' -f 1 | rev)

build:
	#building app static binary ...
	env GOOS=linux CGO_ENABLED=0 go build -a -installsuffix nocgo $(BUILD_FLAGS) -o $(BINARY) ./cmd/.

clean:
#golang linting ...
	golint $(GOPACKAGES)
	go vet $(GOPACKAGES)

coverage:
	#running test coverage ...
	go test -cover $(GOPACKAGES)

generate:
#running go generate tool ...
	go generate $(GOPACKAGES)

image:
	#building docker image ...
	docker build -t $(IMAGE) .

race:
	#tracing race condition ...
	go test -race -v $(GOPACKAGES)

run:
	go run $(CMD)

test:
	#running unit tests. ..
	go test -p 6 -v $(GOPACKAGES)