FROM golang:1.22.4-alpine3.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -installsuffix cgo -o http-service ./cmd/main.go
CMD ["./http-service"]