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