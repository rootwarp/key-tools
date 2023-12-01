.PHONY: build clean

test:
	go test ./... -v

build:
	go build -o key-tool

clean:
	go clean --cache
	rm -rf key-tool
