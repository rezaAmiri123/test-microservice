FROM golang:1.17-alpine

WORKDIR /app

ENV CONFIG=docker

COPY . /app
RUN apk add bash
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod vendor


ENTRYPOINT CompileDaemon --build="go build -o main user_service/cmd/user/main.go" --command=./main



#FROM golang:1.17-buster AS build

#WORKDIR /app
#
#COPY . /app
#RUN go mod download
#
#RUN cd user_service/cmd/user && go build -o /docker-gs-ping && cd  ../../..
#
###
### Deploy
###
##FROM gcr.io/distroless/base-debian10
#FROM ellerbrock/alpine-bash-curl-ssl
#WORKDIR /
#
#COPY --from=build /docker-gs-ping /docker-gs-ping
#
##EXPOSE 8080
##
##USER nonroot:nonroot
#
#ENTRYPOINT ["/docker-gs-ping"]