FROM golang:1.23.9 AS builder

WORKDIR /app
COPY . .

RUN go build -o s3-file-server .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/s3-file-server .

CMD ["./s3-file-server"]
