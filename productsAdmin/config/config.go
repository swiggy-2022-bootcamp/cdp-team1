package config

// Server
const (
	SERVICE_NAME     = "Product Admin Service"
	REST_SERVER_PORT = "9119"
)

// Log
const (
	LOG_FILE      = "./server.log"
	LOG_FILE_MODE = false
)

// AWS DynamoDB
const (
	DYNAMO_DB_REGION = "us-east-1"
	DYNAMO_DB_URL    = "https://dynamodb.us-east-1.amazonaws.com"
)

// gRPC Server config for Get Quantity for ProductID
const (
	GRPC_PRODUCT_QUANTITY_SERVER_IP   = "localhost" //"product-front-store" //
	GRPC_PRODUCT_QUANTITY_SERVER_PORT = "19091"
)

// gRPC Server config for Get Price for ProductID
const (
	GRPC_PRODUCT_PRICE_SERVER_PORT = "19191"
)
