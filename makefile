build:
	go build -o ./artifacts/app ./cmd/main.go

test:
	go test ./internal/...
