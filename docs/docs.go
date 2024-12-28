// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "Retrieve a list of products with optional filters and applied discounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get Products",
                "operationId": "get-products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category to filter by (e.g., boots, sandals, sneakers)",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum price in cents/pennies (e.g., 89000 for 890.00€)",
                        "name": "priceLessThan",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved products",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.getProductsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_tumenbayev_go-products-api_internal_entity.Price": {
            "description": "Price contains pricing information for a product",
            "type": "object",
            "properties": {
                "currency": {
                    "description": "@Description Currency code (e.g., EUR)",
                    "type": "string"
                },
                "discount_percentage": {
                    "description": "@Description Discount percentage if applicable",
                    "type": "string"
                },
                "final": {
                    "description": "@Description Final price after applying discounts in cents/pennies",
                    "type": "integer"
                },
                "original": {
                    "description": "@Description Original price in cents/pennies",
                    "type": "integer"
                }
            }
        },
        "github_com_tumenbayev_go-products-api_internal_entity.Product": {
            "description": "Product represents a product in the system",
            "type": "object",
            "properties": {
                "category": {
                    "description": "@Description Category of the product",
                    "type": "string"
                },
                "name": {
                    "description": "@Description Name of the product",
                    "type": "string"
                },
                "price": {
                    "description": "@Description Price information of the product",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_tumenbayev_go-products-api_internal_entity.Price"
                        }
                    ]
                },
                "sku": {
                    "description": "@Description Unique identifier for the product",
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.getProductsResponse": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_tumenbayev_go-products-api_internal_entity.Product"
                    }
                }
            }
        },
        "internal_controller_http_v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Products API",
	Description:      "A Products API service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
