FROM golang:alpine

RUN apk update && apk add git

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN cd /app && git clone -b Rewards https://github.com/swiggy-2022-bootcamp/cdp-team1.git
WORKDIR /app/cdp-team1/rewards
RUN go build -o main .
EXPOSE 9121
EXPOSE 19071
CMD [ "/app/cdp-team1/rewards/main" ]