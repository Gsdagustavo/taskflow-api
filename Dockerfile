# syntax=docker/dockerfile:1

FROM golang:1.23-alpine AS builder

WORKDIR /taskflow

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# --- final lightweight image ---
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /taskflow/main .

EXPOSE 8080
CMD ["./main"]
