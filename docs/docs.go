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
        "/v1/all": {
            "get": {
                "description": "Посмотреть все доступные песни с данными",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MockServer"
                ],
                "summary": "Get All from mockserver",
                "operationId": "get-all",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/v1/songs": {
            "get": {
                "description": "get songs from library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get songs",
                "operationId": "get-song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by group",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by song",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by created_at",
                        "name": "releaseDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by text",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by link",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "add song to library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Create song",
                "operationId": "create-song",
                "parameters": [
                    {
                        "description": "song info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SongCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/v1/songs/verse/{id}": {
            "get": {
                "description": "get verse of song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Get verse",
                "operationId": "get-verse",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Number of verse (номер куплета)",
                        "name": "num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Verse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/v1/songs/{id}": {
            "put": {
                "description": "Update song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Update song",
                "operationId": "update-song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "song info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SongUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete song from library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Song"
                ],
                "summary": "Delete song",
                "operationId": "delete-song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helpers.Response": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "request": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "helpers.Verse": {
            "type": "object",
            "properties": {
                "verse": {
                    "type": "string",
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?"
                }
            }
        },
        "model.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string",
                    "default": "NOT FOUND",
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "releaseDate": {
                    "type": "string",
                    "default": "NOT FOUND",
                    "example": "16.07.2006"
                },
                "song": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "default": "NOT FOUND",
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
                }
            }
        },
        "model.SongCreate": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "song": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                }
            }
        },
        "model.SongUpdate": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string",
                    "example": "Muse"
                },
                "link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "16.07.2006"
                },
                "song": {
                    "type": "string",
                    "example": "Supermassive Black Hole"
                },
                "text": {
                    "type": "string",
                    "example": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
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
	Title:            "Music library API",
	Description:      "API Server for Music library application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
