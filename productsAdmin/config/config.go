package config

// Server
const (
	SERVICE_NAME = "Product Admin Service"
	SERVER_PORT  = "9119"
)

// Log
const (
	LOG_FILE      = "./server.log"
	LOG_FILE_MODE = false
)

// AWS DynamoDB
const (
	AWS_CREDENTIALS_FILE = "./awsCredentials.txt"
	DYNAMO_DB_REGION     = "us-east-1"
	DYNAMO_DB_URL        = "https://dynamodb.us-east-1.amazonaws.com"
)
