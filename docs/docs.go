// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "email": "fiber@swagger.io"
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
        "/api/task/createTask": {
            "post": {
                "description": "create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to set task params",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/task/deliteTask": {
            "delete": {
                "description": "delete task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to delete task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/task/getUser": {
            "get": {
                "description": "gets task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for finding task by id",
                        "name": "task_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding task by name",
                        "name": "task_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding task by desc",
                        "name": "task_description",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "task data",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/task/updateTask": {
            "patch": {
                "description": "update task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to update task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/track/createTrack": {
            "post": {
                "description": "create track",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to set track",
                        "name": "track",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/track/deliteTrack": {
            "delete": {
                "description": "delete track by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to delete track",
                        "name": "track",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/track/getTrack": {
            "get": {
                "description": "gets track",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for finding track by id",
                        "name": "track_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "used for finding track by user id",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "used for finding track by task id",
                        "name": "task_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "track data",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/track/updateTrack": {
            "patch": {
                "description": "update track by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to update track",
                        "name": "track",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/createUser": {
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to set user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/deleteUser": {
            "delete": {
                "description": "delete user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to delete user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/getUser": {
            "get": {
                "description": "gets user by id/name/surname/...",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for finding user by id",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding user by name",
                        "name": "user_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding user by surname",
                        "name": "user_surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding user by patronymic",
                        "name": "user_patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "used for finding user by address",
                        "name": "user_address",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "used for finding user by passport",
                        "name": "user_passport",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user data",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/getUsers": {
            "get": {
                "description": "getting users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "max items on page",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "starts from 1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "list of users",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/labourCosts": {
            "get": {
                "description": "gives tracks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "used for finding tracks",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "used for left edge of dates",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "used for right edge of dates",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "list of tracks",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        },
        "/api/user/updateUser": {
            "patch": {
                "description": "update user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "used to update user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/handler.Output"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Output": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                }
            }
        },
        "model.TaskCreate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "code review by lead"
                },
                "name": {
                    "type": "string",
                    "example": "code review"
                }
            }
        },
        "model.TaskDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.TaskUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "code review by lead"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "code review"
                }
            }
        },
        "model.TrackCreate": {
            "type": "object",
            "properties": {
                "start": {
                    "type": "string",
                    "example": "2006-01-02 15:04:05"
                },
                "stop": {
                    "type": "string",
                    "example": "2010-01-02 15:04:05"
                },
                "taskID": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.TrackDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.TrackUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "start": {
                    "type": "string",
                    "example": "2006-01-02 15:04:05"
                },
                "stop": {
                    "type": "string",
                    "example": "2010-01-02 15:04:05"
                },
                "taskID": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.UserCreate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "LA"
                },
                "name": {
                    "type": "string",
                    "example": "Jon"
                },
                "passport": {
                    "type": "integer",
                    "example": 123456
                },
                "patronymic": {
                    "type": "string",
                    "example": "White"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                }
            }
        },
        "model.UserDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.UserUpdate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "LA"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Jon"
                },
                "passport": {
                    "type": "integer",
                    "example": 123456
                },
                "patronymic": {
                    "type": "string",
                    "example": "White"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
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
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
