##
## Build
##

FROM golang:1.16-alpine

#RUN mkdir /app
WORKDIR /app

ADD . /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build -o main .

CMD ["/app/main"]