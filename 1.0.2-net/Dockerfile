FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o main .

EXPOSE 8080 

RUN apk update && apk --no-cache add bash curl

ENV HELLO_MESSAGE="Hello, Universe!"
ENV PORT=8080

ENV MONGO_ADDR=mongo
ENV MONGO_PORT=27017
ENV REDIS_ADDR=redis
ENV REDIS_PORT=6379

COPY main.go /app/main.go

CMD ["./main"]