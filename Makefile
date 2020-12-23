GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build

all: test build
build:
	rm -rf dist/
	$(GOBUILD) -o dist/bin/replaytest main.go

test:
	$(GOTEST) -v ./...

clean:
	rm -rf dist/

run:
	dist/bin/replaytest

stop:
	pkill -f dist/bin/cmd