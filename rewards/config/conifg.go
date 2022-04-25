package config

// Server
const (
	SERVICE_NAME = "Reward Service"
	SERVER_PORT  = "9121"
)

// Log
const (
	LOG_FILE      = "./server.log"
	LOG_FILE_MODE = true
)

// AWS DynamoDB
const (
	AWS_CREDENTIALS_FILE = "./awsCredentials.txt"
	DYNAMO_DB_REGION     = "us-east-1"
	DYNAMO_DB_URL        = "https://dynamodb.us-east-1.amazonaws.com"
)

// gRPC Server config for reward points
const (
	GRPC_SERVER_IP   = "localhost" //"Trabnsaction" //
	GRPC_SERVER_PORT = "19071"
)
