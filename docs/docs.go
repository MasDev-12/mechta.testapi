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
        "/url/shortener": {
            "post": {
                "description": "Create a shortened URL by passing the URL data in the request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URLs"
                ],
                "summary": "Create a new shortened URL",
                "parameters": [
                    {
                        "description": "URL data",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "URL successfully shortened",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateUrlResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid URL data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request timed out",
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
        },
        "/url/shortener/{userId}": {
            "get": {
                "description": "Get all shortened URLs by providing a user ID in the request URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URLs"
                ],
                "summary": "Get all URLs created by a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved user URLs",
                        "schema": {
                            "$ref": "#/definitions/responses.GetUserUrlsResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No URLs found for the user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request timed out",
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
        },
        "/url/stats/{link}": {
            "get": {
                "description": "Get statistics for a URL using its short name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URLs"
                ],
                "summary": "Get URL statistics by short name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL to get statistics for",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved URL statistics",
                        "schema": {
                            "$ref": "#/definitions/responses.GetUrlStatByShortNameResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid link",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "URL not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request timed out",
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
        },
        "/url/{link}": {
            "get": {
                "description": "Retrieve the original URL by providing the short URL in the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URLs"
                ],
                "summary": "Get original URL by short name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved the original URL",
                        "schema": {
                            "$ref": "#/definitions/responses.GetUrlByShortNameResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid link",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "URL not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request timed out",
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
            },
            "delete": {
                "description": "DeleteByShortName a URL by providing its short name in the request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URLs"
                ],
                "summary": "DeleteByShortName URL by short name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL to delete",
                        "name": "link",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted the URL",
                        "schema": {
                            "$ref": "#/definitions/responses.DeleteUrlByShortNameResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid link",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "URL not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request timed out",
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
        },
        "/user/create": {
            "post": {
                "description": "Create a new user by passing user data in the request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully created",
                        "schema": {
                            "$ref": "#/definitions/responses.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
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
        },
        "/user/{id}": {
            "get": {
                "description": "Get a user by passing the user ID in the query parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/responses.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "408": {
                        "description": "Request Timeout",
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
        "dto.URLDto": {
            "type": "object",
            "properties": {
                "click_count": {
                    "type": "integer"
                },
                "expires_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_accessed_at": {
                    "description": "Опускается, если nil",
                    "type": "string"
                },
                "original_url": {
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                },
                "user_id": {
                    "description": "Внешний ключ на пользователя",
                    "type": "string"
                }
            }
        },
        "requests.CreateURLRequest": {
            "type": "object",
            "required": [
                "original_url",
                "user_id"
            ],
            "properties": {
                "original_url": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "requests.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "Валидация для корректного email",
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
        "responses.CreateUrlResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "id": {
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                }
            }
        },
        "responses.CreateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "error": {},
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.DeleteUrlByShortNameResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {
                    "type": "boolean"
                }
            }
        },
        "responses.GetUrlByShortNameResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "url": {
                    "$ref": "#/definitions/dto.URLDto"
                }
            }
        },
        "responses.GetUrlStatByShortNameResponse": {
            "type": "object",
            "properties": {
                "click_count": {
                    "type": "integer"
                },
                "error": {},
                "expires_at": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_accessed_at": {
                    "type": "string"
                },
                "original_url": {
                    "type": "string"
                }
            }
        },
        "responses.GetUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "error": {},
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.GetUserUrlsResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "urls": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.URLDto"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
