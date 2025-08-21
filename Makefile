BINARY_NAME=dupclass

build:
	go build -ldflags="-s -w" -gcflags="-B -C -l=4" -a -o $(BINARY_NAME) ./src

test:
	go test -v ./src

# Generate big test files for performance testing
generate-stress-test:
	cd testdata/generator && go run generate_test_files.go custom 10000 20

# Run stress test with generated files
stress-test: build generate-stress-test
	time ./$(BINARY_NAME) testdata/generated/test_files_custom_10000_20kb

clean:
	rm -f $(BINARY_NAME)
	rm -rf testdata/generated/test_files_*

.PHONY: build test generate-stress-test stress-test clean