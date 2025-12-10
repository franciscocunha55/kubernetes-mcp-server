FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o kubernetes-mcp-server .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/kubernetes-mcp-server .

EXPOSE 8080

CMD ["./kubernetes-mcp-server", "http"]