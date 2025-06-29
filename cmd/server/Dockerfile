# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server ./cmd/server

# Final stage
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/grpc-server .
COPY --from=builder /app/db/migrations ./db/migrations
USER nonroot:nonroot
EXPOSE 50051
ENTRYPOINT ["/grpc-server"]