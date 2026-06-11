FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o mail-service ./cmd/...

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/mail-service .
EXPOSE 8091
ENTRYPOINT ["/app/mail-service"]
