.PHONY: run lint

run:
	go run ./cmd/rk_sensor/main.go

lint:
	gofmt -s -w .
	golangci-lint run --config .golangci.yml