# FROM golang:1.17-alpine

# WORKDIR /app

# ENV CONFIG=docker

# COPY . /app
# RUN apk add bash
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go mod vendor


# ENTRYPOINT CompileDaemon --build="go build -o main library_service/cmd/library/main.go" --command=./main


FROM golang:1.17-buster AS build

WORKDIR /app

COPY . /app
RUN go mod download

RUN CGO_ENABLED=0 go build -o main library_service/cmd/library/main.go
## Deploy
##
#FROM gcr.io/distroless/base-debian10
FROM ellerbrock/alpine-bash-curl-ssl
WORKDIR /

COPY --from=build /app/main /main

#EXPOSE 8080
#
#USER nonroot:nonroot

ENTRYPOINT ["/main"]
