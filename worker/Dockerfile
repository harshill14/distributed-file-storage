# Build Stage
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o worker .

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/worker .
EXPOSE 50052
CMD ["./worker"]
