##
## Build
##

FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build -o main .

EXPOSE 7000

CMD ["/app/main"]