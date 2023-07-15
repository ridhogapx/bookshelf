FROM golang:1.20.6-alpine3.18

WORKDIR /bookshelf

COPY go.mod /bookshelf

RUN go mod download

COPY . /bookshelf 

RUN go mod tidy

RUN go build -o /bookshelf/bin/server /bookshelf/server/main.go

EXPOSE 50051

CMD ["/bookshelf/bin/server"]