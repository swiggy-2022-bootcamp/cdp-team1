##
## Build
##

FROM golang:1.16-alpine as build

#RUN mkdir /app
WORKDIR /app

ADD . /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build -o main .
COPY .env .

FROM alpine

COPY --from=build /app/main /.
COPY --from=build /app/.env /.

CMD ["/main"]