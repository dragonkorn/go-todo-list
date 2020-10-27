
FROM golang:latest

WORKDIR /src

COPY ./ /src

RUN go mod download

ENTRYPOINT go run src/main.go