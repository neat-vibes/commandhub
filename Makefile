BINARY_NAME=commandhub

build:
	@ printf "Building $(BINARY_NAME)...\n"
	go build -o ./bin/$(BINARY_NAME) -v ./cmd/
	@ printf "Done.\n"
