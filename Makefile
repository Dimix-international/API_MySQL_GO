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

swagger_install:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	rmdir /S /Q docs && mkdir docs && swag init -g ./cmd/main.go -dir .	

swagger-ubu:
	rm -rf docs && mkdir docs && swag init -g ./cmd/main.go -dir .	