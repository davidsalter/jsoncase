# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=jsoncase
SRC_DIR=.

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) $(SRC_DIR)

# Run the project
run:
	$(GORUN) $(SRC_DIR)/jsoncase.go example/input.json

# Test the project
test:
	$(GOTEST) -v ./...

# Clean the build
clean:
	rm -f $(BINARY_NAME)

# Help
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the project"
	@echo "  run     - Run the project"
	@echo "  test    - Test the project"
	@echo "  clean   - Clean the build"
	@echo "  help    - Show this help message"

.PHONY: build run test clean help