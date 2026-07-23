FROM golang:1.22-alpine AS builder
WORKDIR /application
COPY . .
RUN go mod init application
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM scratch
WORKDIR /application
COPY --from=builder /application/app .
EXPOSE 8080
CMD ["./app"]
