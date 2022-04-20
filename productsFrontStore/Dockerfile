FROM golang:alpine

RUN apk update && apk add git

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN cd /app && git clone -b products-front-store https://github.com/swiggy-2022-bootcamp/cdp-team1.git
WORKDIR /app/cdp-team1/productsFrontStore
RUN go build -o main .

EXPOSE 9009
EXPOSE 19091

CMD [ "/app/cdp-team1/productsFrontStore/main" ]