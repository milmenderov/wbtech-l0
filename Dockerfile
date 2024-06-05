FROM golang:1.21.5-alpine3.20
WORKDIR /app
COPY go.mod go.sum ./
COPY . .

RUN go build -a -installsuffix cgo -o httpservice ./cmd/main.go
CMD ["./httpservice"]