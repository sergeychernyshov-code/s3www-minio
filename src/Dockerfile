FROM golang:1.24.3 AS builder

WORKDIR /app
COPY . .

# Disable CGO to make static binary
ENV CGO_ENABLED=0

RUN go build -o s3-file-server .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/s3-file-server .

CMD ["./s3-file-server"]
