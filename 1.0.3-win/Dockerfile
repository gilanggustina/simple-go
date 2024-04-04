FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy  your go source code (replace with your actual file name)
COPY main.go .

ENV GOOS=windows GOARCH=amd64

RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD [ "/app/main" ]

