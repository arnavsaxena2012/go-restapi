# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go get -u github.com/gorilla/mux

RUN go mod download

RUN go build -o main .

EXPOSE 9090

CMD ["./main"]
