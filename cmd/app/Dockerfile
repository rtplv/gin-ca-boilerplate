FROM golang:alpine

WORKDIR /go/src/example
ENV TZ=Europe/Moscow

ADD . .

# For change tz
RUN apk add --no-cache tzdata

EXPOSE 3000

CMD go build -o ./bin/server ./cmd/app/ &&\
    go get -u github.com/swaggo/swag/cmd/swag &&\
    swag init -g ./cmd/app/main.go &&\
    ./bin/server
