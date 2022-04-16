FROM golang:alpine

RUN apk update && apk add git

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN cd /app && git clone -b products-admin https://github.com/swiggy-2022-bootcamp/cdp-team1.git
WORKDIR /app/cdp-team1/productsAdmin
RUN go build -o main .

CMD [ "/app/cdp-team1/productsAdmin/main" ]