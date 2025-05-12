FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o s3-file-server .

FROM gcr.io/distroless/static
COPY --from=builder /app/s3-file-server /s3-file-server
ENTRYPOINT ["/s3-file-server"]