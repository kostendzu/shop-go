package docs

import (
	"currency/pkg/config"
	"log"
	"os"

	"github.com/swaggo/swag"
)

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "description": "API для получения курсов валют",
        "title": "Currency API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "schemes": [
        "http"
    ],
    "paths": {
        "/currencies": {
            "get": {
                "tags": ["Currencies"],
                "summary": "Get currencies",
                "description": "Get all currencies if no date parameter is provided. If a date parameter is provided, get currencies for the specified date.",
                "operationId": "getCurrencies",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "name": "date",
                        "in": "query",
                        "description": "Date in the format YYYY-MM-DD",
                        "required": false,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/CurrencyRate"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request (invalid date format)",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CurrencyRate": {
            "type": "object",
            "properties": {
                "cur_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string",
                    "format": "date-time"
                },
                "cur_abbreviation": {
                    "type": "string"
                },
                "cur_scale": {
                    "type": "integer"
                },
                "cur_name": {
                    "type": "string"
                },
                "cur_officialRate": {
                    "type": "number",
                    "format": "float"
                }
            }
        }
    }
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Currency API",
	Description:      "API для получения курсов валют",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	nodeEnv := os.Getenv("NODE_ENV")

	log.Println("ne: ", nodeEnv)
	if nodeEnv != "DOCKER" {
		err := config.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not set in the environment variables")
	}
	SwaggerInfo.Host = "localhost:" + port

	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
