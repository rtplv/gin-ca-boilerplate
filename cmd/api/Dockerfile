FROM golang:alpine

WORKDIR /go/src/example
ENV TZ=Europe/Moscow

ADD . .

# For change tz
RUN apk add --no-cache tzdata

EXPOSE 3000

CMD go build -o ./bin/server ./cmd/api/ &&\
    go get -u github.com/swaggo/swag/cmd/swag &&\
    swag init -g ./cmd/api/main.go &&\
    ./bin/server
