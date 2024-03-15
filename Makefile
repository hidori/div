.PHONY: install
install:
	go mod download

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	time go run ./cmd/main.go
