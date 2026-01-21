# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum from services directory
COPY services/go.mod services/go.sum ./
RUN go mod download

# Copy the entire services directory
COPY services/ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# ---- Runtime stage ----
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
