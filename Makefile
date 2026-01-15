# Variables
BINARY_NAME=crawler
DEFAULT_URL=https://crawler-test.com/

test:
	go test ./...

testv:
	go test -v ./...

build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) .

# Run "time" on the binary with optional arguments
time: build
	@echo "------------------------------------------------"
	@echo "Running: ./$(BINARY_NAME) $(or $(URL),$(DEFAULT_URL)) $(CONCURRENCY) $(MAX_PAGES)"
	@echo "------------------------------------------------"
	@time ./$(BINARY_NAME) $(or $(URL),$(DEFAULT_URL)) $(CONCURRENCY) $(MAX_PAGES)

clean:
	rm -f $(BINARY_NAME)

.PHONY: build time clean