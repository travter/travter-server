FROM golang:1.19.2-alpine3.16

WORKDIR /server

COPY ../.. .

RUN go build -o bin/build cmd/main.go

CMD ["bin/build"]