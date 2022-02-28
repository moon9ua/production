// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.economicus.kr/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "이메일, 비밀번호로 로그인하기",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Local login",
                "parameters": [
                    {
                        "description": "User login email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "User login password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.Token"
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/quants": {
            "get": {
                "description": "메인 화면에서 사용될 퀀트 모델들 반환, 개선 필요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Return a list of quant models",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "number of page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "number of elements",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "\"\"",
                        "description": "fields for order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "\"\"",
                        "description": "keyword for query",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of quants",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Quant"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/quants/quant": {
            "post": {
                "description": "실험실에서 모델 만들기를 눌렀을 때, 모델 생성",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Create a quant model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Quant option data",
                        "name": "quant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.QuantC"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/quants/quant-option/{quant_id}": {
            "patch": {
                "description": "퀀트 옵션을 변경하고자 할 경우 사용",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Update a quant option",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Quant option data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.QuantOptU"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "QuantID to update",
                        "name": "quant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/quants/quant/{quant_id}": {
            "get": {
                "description": "모델 상세페이지에서 사용될 퀀트 모델 반환, 개선 필요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Return a quant model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID of a quant",
                        "name": "quant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A quant",
                        "schema": {
                            "$ref": "#/definitions/model.Quant"
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            },
            "delete": {
                "description": "퀀트 모델을 제거할 경우 사용",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Delete a quant model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quant ID to delete",
                        "name": "quant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            },
            "patch": {
                "description": "모델 저장버튼(activate)을 누르거나, 모델 설명을 변경하고자 할 경우 사용",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quant"
                ],
                "summary": "Update a quant model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Quant data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.QuantC"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/refresh-token": {
            "post": {
                "description": "Access token 기간 만료시, Refresh token을 사용하여 jwt 토큰 재발급",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh jwt token",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "refreshed access token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "이메일, 비밀번호로 유저 회원가입",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register a user",
                "parameters": [
                    {
                        "description": "A user information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "이메일, 비밀번호로 유저 회원가입",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Return all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer {access_token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "A user information",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "404": {
                        "description": "Not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.httpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.httpError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Some error message"
                }
            }
        },
        "jwt.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "model.Quant": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "quant model description"
                },
                "name": {
                    "type": "string",
                    "example": "quant model name"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "quant.DoublePair": {
            "type": "object",
            "properties": {
                "max": {
                    "type": "number",
                    "example": -100.01
                },
                "min": {
                    "type": "number",
                    "example": 100.01
                }
            }
        },
        "quant.IntPair": {
            "type": "object",
            "properties": {
                "max": {
                    "type": "integer",
                    "example": -10000
                },
                "min": {
                    "type": "integer",
                    "example": 1000000
                }
            }
        },
        "request.QuantC": {
            "type": "object",
            "properties": {
                "de_ratio": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "dividend_payout_ratio": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "dividend_yield": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "end_date": {
                    "type": "string",
                    "example": "2021-03-31T00:00:000.Z"
                },
                "financing": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "investing": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "main_sectors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "IT",
                        "소재"
                    ]
                },
                "market_cap": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "name": {
                    "type": "string",
                    "example": "Model name"
                },
                "net_profit": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "net_profit_rate": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "net_revenue": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "net_revenue_rate": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "operating": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "pbr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "pcr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "per": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "psr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "roa": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "roe": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "start_date": {
                    "type": "string",
                    "example": "2016-03-31T00:00:000.Z"
                }
            }
        },
        "request.QuantOptU": {
            "type": "object",
            "properties": {
                "de_ratio": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "dividend_payout_ratio": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "dividend_yield": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "end_date": {
                    "type": "string",
                    "example": "2021-03-31T00:00:000.Z"
                },
                "financing": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "investing": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "main_sectors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "IT",
                        "소재"
                    ]
                },
                "market_cap": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "net_profit": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "net_profit_rate": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "net_revenue": {
                    "$ref": "#/definitions/quant.IntPair"
                },
                "net_revenue_rate": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "operating": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "pbr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "pcr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "per": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "psr": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "roa": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "roe": {
                    "$ref": "#/definitions/quant.DoublePair"
                },
                "start_date": {
                    "type": "string",
                    "example": "2016-03-31T00:00:000.Z"
                }
            }
        },
        "request.RegisterRequest": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "string",
                    "example": "2016-03-31T00:00:000.Z"
                },
                "email": {
                    "type": "string",
                    "example": "example@economicus.kr"
                },
                "name": {
                    "type": "string",
                    "example": "user name"
                },
                "nickname": {
                    "type": "string",
                    "example": "user nickname"
                },
                "password": {
                    "type": "string",
                    "example": "some password"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{"https", "http"},
	Title:            "Economicus Main Server",
	Description:      "Economicus 메인 서버",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
