// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Ayan Dutta",
            "email": "ayan59dutta@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "get": {
                "description": "Heath Check",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "responses": {
                    "200": {
                        "description": "Service is healthy",
                        "schema": {
                            "$ref": "#/definitions/dto.HealthResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Service is unavailable",
                        "schema": {
                            "$ref": "#/definitions/dto.HealthResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "User login",
                        "name": "dto.LoginRequestDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LoginResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.LogoutResponseDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Verify Token",
                        "name": "dto.VerificationRequestDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerificationRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.HealthResponseDTO": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequestDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponseDTO": {
            "type": "object",
            "properties": {
                "auth_token": {
                    "type": "string"
                }
            }
        },
        "dto.LogoutResponseDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.VerificationRequestDTO": {
            "type": "object",
            "properties": {
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "errs.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer Token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Auth Service APIs",
	Description:      "Swagger API for Golang Project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
