FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=builder /backend /backend
EXPOSE 8080 9999
ENTRYPOINT ["/backend"]
