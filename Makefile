BINARY_NAME=dupclass

build:
	go build -ldflags="-s -w" -gcflags="-B -C -l=4" -a -o $(BINARY_NAME) ./src

test:
	go test -v ./src

clean:
	rm -f $(BINARY_NAME)

.PHONY: build test clean