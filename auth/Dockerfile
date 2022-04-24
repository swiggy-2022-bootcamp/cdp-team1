FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git build-base && apk add tzdata
ENV TZ=Asia/Kolkata

# Set destination for COPY
WORKDIR /auth

# Copy the source code.
COPY  . .

# get all the dependencies
RUN go get ./...

#WORKDIR /

# Build
RUN go build -o /authService

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8081

CMD [ "/authService" ]