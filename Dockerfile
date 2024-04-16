# Build the application from source
FROM golang:1.22.0-alpine

WORKDIR /app/src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ../API_MySQL_GO ./cmd/main.go

WORKDIR /app

RUN rm -rf src

EXPOSE 8080

#Run
CMD ["/app/API_MySQL_GO"]