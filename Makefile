run:
	go run ./cmd/main.go

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./...

mockery-install:
	go get github.com/vektra/mockery/v2@v2.38.0

generate-mocks:
	go generate ./...