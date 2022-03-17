// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "get": {
                "description": "returns user personal data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "create new user on system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserCreateArgs"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.NoContentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "mark user as deleted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.NoContentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    }
                }
            }
        },
        "/user/authorization": {
            "post": {
                "description": "authorization user on system by email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "authorization",
                "parameters": [
                    {
                        "description": "User auth data",
                        "name": "authorization",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserAuthorizationArgs"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.UserAuthorizationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "post": {
                "description": "edit user password: add new or replace old password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "edit password",
                "parameters": [
                    {
                        "description": "New password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserPasswordEditArgs"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.NoContentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/liberror.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.NoContentResponse": {
            "type": "object"
        },
        "handlers.ProfileResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "user create datetime",
                    "type": "string",
                    "example": "0001-01-01T00:00:00Z"
                },
                "email": {
                    "description": "email",
                    "type": "string",
                    "example": "some@mail.com"
                },
                "login": {
                    "description": "login",
                    "type": "string",
                    "example": "some-login"
                },
                "name": {
                    "description": "name",
                    "type": "string",
                    "example": "some-name"
                },
                "phone": {
                    "description": "phone",
                    "type": "string",
                    "example": "88009998889988"
                }
            }
        },
        "handlers.UserAuthorizationArgs": {
            "type": "object",
            "properties": {
                "login": {
                    "description": "login",
                    "type": "string",
                    "example": "some login"
                },
                "password": {
                    "description": "password",
                    "type": "string",
                    "example": "some_password"
                }
            }
        },
        "handlers.UserAuthorizationResponse": {
            "type": "object",
            "properties": {
                "expired_at": {
                    "description": "token deadline date time",
                    "type": "integer"
                },
                "token": {
                    "description": "access token string",
                    "type": "string"
                }
            }
        },
        "handlers.UserCreateArgs": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "email",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 5,
                    "example": "some@mail.com"
                },
                "login": {
                    "description": "login",
                    "type": "string",
                    "example": "somelogin"
                },
                "name": {
                    "description": "name",
                    "type": "string",
                    "example": "some name"
                },
                "password": {
                    "description": "password",
                    "type": "string",
                    "maxLength": 1024,
                    "minLength": 7,
                    "example": "secret-word"
                },
                "phone": {
                    "description": "phone",
                    "type": "string",
                    "example": "88009998889999"
                }
            }
        },
        "handlers.UserPasswordEditArgs": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "new password",
                    "type": "string",
                    "example": "some_password"
                }
            }
        },
        "liberror.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Growth-place API",
	Description: "This is a growth service for managing personal targets",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
