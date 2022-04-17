FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

#Copy the source code.
COPY . .

# get all the dependencies
RUN go get ./...

# build
RUN go build ./main.go

EXPOSE 9003

COPY .env  .

CMD [ "./main" ] 