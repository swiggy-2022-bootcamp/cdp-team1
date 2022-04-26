FROM golang:1.18-alpine as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code.
COPY  . .

# Build
RUN go build -o /paymentService

COPY .env .

FROM alpine

COPY --from=build /paymentService /.
COPY --from=build /app/.env /.

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 9000
EXPOSE 9004

# Run
ENTRYPOINT ["/paymentService"]