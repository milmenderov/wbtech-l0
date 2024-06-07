FROM golang:1.21.4-alpine3.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -installsuffix cgo -o http-service ./cmd/main.go
CMD ["./http-service"]