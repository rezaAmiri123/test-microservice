FROM golang:1.17-buster AS build

WORKDIR /app

COPY . /app
RUN go mod download

RUN CGO_ENABLED=0 go build -o main api_service/cmd/api/main.go
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
