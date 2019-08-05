BINARY_NAME=deployer
all:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME) -v