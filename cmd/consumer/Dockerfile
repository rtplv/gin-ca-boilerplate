FROM golang:alpine

WORKDIR /go/src/example-consumer
ENV TZ=Europe/Moscow

ADD . .

# For change tz
RUN apk add --no-cache tzdata

CMD go build -o ./bin/consumer ./cmd/consumer/ &&\
    ./bin/consumer
