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
        "/lyrics": {
            "get": {
                "description": "Returns lyrics for a specific song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lyrics"
                ],
                "summary": "Get song lyrics",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "songid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit for pagination",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Lyrics"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid arguments",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Wrong method",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error processing request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song": {
            "post": {
                "description": "Add a new song and its lyrics to the database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Create a new song",
                "parameters": [
                    {
                        "description": "Song information",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.NewSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Song added successfully"
                    },
                    "400": {
                        "description": "Bad request",
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
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a song by its ID from the database.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Delete song by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID to delete",
                        "name": "songid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "Method not allowed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing song in the library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Update a song",
                "parameters": [
                    {
                        "description": "Song update info",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateSongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Bad request",
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
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Gets songs by name, id, release date, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Gets songs by specified conditions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "songid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Song Name",
                        "name": "songname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "groupid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Group Name",
                        "name": "groupname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Release date start in YYYY-MM-DD format",
                        "name": "releasedate-start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Release date end in YYYY-MM-DD format",
                        "name": "releasedate-end",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit for pagination",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Lyrics": {
            "type": "object",
            "properties": {
                "lyrics": {
                    "type": "string"
                },
                "orderN": {
                    "type": "integer"
                },
                "songId": {
                    "type": "integer"
                }
            }
        },
        "models.NewSongRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "models.Song": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "groupName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                }
            }
        },
        "models.UpdateSongRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3030",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Song Library API",
	Description:      "API for managing a song library",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
