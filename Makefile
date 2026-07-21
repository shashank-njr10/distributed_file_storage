build:
	@go build -o bin/main main.go

run: build
	@./bin/main

test:
	@go test ./... -v


# '@' prevents the command from being printed to the terminal when executed.