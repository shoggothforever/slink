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
        "contact": {
            "name": "Kalun",
            "url": "124.220.190.203",
            "email": "1337231450@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/url/create": {
            "post": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Url"
                ],
                "summary": "创建短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "long url",
                        "name": "origin",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "short url",
                        "name": "short",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "comment",
                        "name": "comment",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/url/delete": {
            "delete": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Url"
                ],
                "summary": "删除短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "urlId",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DeleteResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/url/pause": {
            "post": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Url"
                ],
                "summary": "冻结短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "urlId",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/url/query": {
            "post": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Url"
                ],
                "summary": "查询短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "urlId",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.QueryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/url/update": {
            "put": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Url"
                ],
                "summary": "更新短链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "urlId",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "short url",
                        "name": "short",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "comment",
                        "name": "comment",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateResponse"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "login with correct account and password",
                "consumes": [
                    "*/*",
                    "application/json"
                ],
                "produces": [
                    "*/*",
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "psw",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "description": "logout",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户退出",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/record/get": {
            "get": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户登录信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.InfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register with account,email and password",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "注册用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user/url/get": {
            "get": {
                "description": "get user's information",
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户url信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwt",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.QueryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DeleteResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.InfoResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "jwt": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.QueryResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "url": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UrlInfo"
                    }
                }
            }
        },
        "model.RegisterResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.UpdateResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.UrlInfo": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "expireTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "origin": {
                    "type": "string"
                },
                "short": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UrlInfo"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "shortLink API",
	Description:      "This is a sample shortLink server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
