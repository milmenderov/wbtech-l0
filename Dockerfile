FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /app

ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN chown -R ${USER}: /app
RUN go build -a -installsuffix cgo -o http-service ./cmd/main.go


FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

COPY --from=builder /app/http-service ./

USER appuser:appuser
CMD ["./http-service"]