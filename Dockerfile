
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

FROM alpine:3.14
RUN apk add --no-cache ca-certificates=20211220-r0
WORKDIR /app
COPY --from=builder /app/myapp .
RUN chmod +x myapp
EXPOSE 4040
CMD ["./myapp"]