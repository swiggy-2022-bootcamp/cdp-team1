# Order service

# Responsibilities
1) Allows user to create, read, update and delete any and all orders from frontstore .
2) Allows user to query orders from CustomerId, OrderId and Status
3) Allows user to update order status and create order invoice

## Technologies Used
- Golang
- Gin Gonic
- gRPC
- DynamoDB
- Docker & Docker Compose
- Jenkins
- Swagger UI

## Features implemented
- [x] Dockerized the service
- [x] Swagger Documentation
- [x] Effective Use of GRPC for inter-service communication.
- [x] Sonarqube to calculate code coverage.
- [x] Implemented REST endpoints using DynamoDB
- [x] Maintained the code repository
- [x] Built the services using CI/CD pipeline
- [x] Data models properly showcased
- [x] Mock Testing using mockgen and testify
- [x] Logging in a separate file
- [x] HealthCheck API

# REST endpoints exposed
1) [POST] - CreateOrder - http://localhost:7000/api/order
2) [POST] - CreateInvoice - http://localhost:7000/api/orders/invoice/:id
3) [GET] - GetAllOrders - http://localhost:7000/api/orders
4) [GET] - GetOrderByStatus- http://localhost:7000/api/orders/status/:status
5) [GET] - GetOrderById - http://localhost:7000/api/orders/:id
6) [GET] - GetOrderByCustomerId - http://localhost:7000/api/orders/users/:id
7) [PUT] - UpdateOrder - http://localhost:7000/api/orders/:id
8) [DELETE] - DeleteOrderById- http://localhost:7000/api/orders/:id
9) [GET] - HealthCheck - http://localhost:7000/api/

# Steps to run application
1) Using docker
   1) 
      `docker build --tag order -t order .`
   2) `docker run -d -p 7000:7000 order ` 

2) Run locally
   1) `cd order` 
   2) `go run .`
