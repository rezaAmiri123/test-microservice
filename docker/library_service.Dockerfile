FROM golang:1.17-alpine

WORKDIR /app

ENV CONFIG=docker

COPY . /app
RUN apk add bash
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod vendor


ENTRYPOINT CompileDaemon --build="go build -o main library_service/cmd/library/main.go" --command=./main
