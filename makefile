build:
	@go build -o totion ./cmd/totion/main.go
run: build
	@./totion