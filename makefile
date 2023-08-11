.PHONY: default run build test docs clean

APP_NAME=gopportunies

default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build: 
	@go build -o $(PP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -f ./docs

