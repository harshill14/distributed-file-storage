# Build Stage
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 50051
CMD ["./server"]
