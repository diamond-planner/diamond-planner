FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o pocketbase

FROM ubuntu:latest
WORKDIR /app
RUN apt update
RUN apt install -y ca-certificates
COPY --from=builder /app/pocketbase ./
COPY --from=builder /app/pb_public ./pb_public

EXPOSE 8090

CMD ["./pocketbase", "serve", "--http=0.0.0.0:8090"]
