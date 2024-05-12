.PHONY: default run build test docs clean

# Variables
APP_NAME=lab-golang-api-crud

# Tasks
default: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -r $(APP_NAME)
	@rm -rf ./docs