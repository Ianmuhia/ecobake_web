
VERSION 0.6
FROM golang:1.19-alpine
WORKDIR /go-example

build:
    COPY go.mod go.sum .
    RUN go mod download
    COPY . .
    RUN mkdir -p "./temp"
    RUN go build -ldflags="-s -w" -o ./temp/ ./...
    SAVE ARTIFACT temp /temp AS LOCAL temp

docker:
    COPY +build/temp .
    ENTRYPOINT ["/go-example/cmd"]
    SAVE IMAGE go-example:latest