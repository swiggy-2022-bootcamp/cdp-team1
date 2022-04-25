FROM golang:1.18-alpine as build

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
RUN go build ./main.go -o shippingService

COPY .env .

FROM alpine

COPY --from=build /shippingService /.
COPY --from=build /app/.env /.

EXPOSE 9003

ENTRYPOINT [ "./main" ]
