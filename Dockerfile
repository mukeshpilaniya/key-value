# Build stage
FROM  golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN  go mod tidy -go=1.18
RUN go build -o main cmd/*

    # Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]