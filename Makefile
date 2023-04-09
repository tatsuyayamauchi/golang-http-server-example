BINARY=guacamole
SOURCES=$(wildcard *.go)
TESTS=$(wildcard *_test.go)

all: clean test build

.PHONY: build
build: $(SOURCES)
	go build -o bin/$(BINARY) guacamole/pkg/cmd/guacamole/main.go

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: run
run: build
	./bin/$(BINARY)
