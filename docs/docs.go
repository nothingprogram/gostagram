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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/instagram/child": {
            "get": {
                "description": "Instagram get child post by parrent id",
                "produces": [
                    "application/json"
                ],
                "summary": "Instagram get child post by parrent id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access token",
                        "name": "access_token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "child id",
                        "name": "child_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.InstaChildResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/instagram/posts": {
            "get": {
                "description": "Instagram get posts by user",
                "produces": [
                    "application/json"
                ],
                "summary": "Instagram get posts by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access token",
                        "name": "access_token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.InstaResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Child": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "media_type": {
                    "type": "string"
                },
                "media_url": {
                    "type": "string"
                },
                "permalink": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                }
            }
        },
        "domain.Children": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Child"
                    }
                }
            }
        },
        "domain.Instagram": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "children": {
                    "$ref": "#/definitions/domain.Children"
                },
                "id": {
                    "type": "string"
                },
                "media_type": {
                    "type": "string"
                },
                "media_url": {
                    "type": "string"
                },
                "permalink": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.InstaChildResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "media_type": {
                    "type": "string"
                },
                "media_url": {
                    "type": "string"
                },
                "permalink": {
                    "type": "string"
                },
                "thumbnail_url": {
                    "type": "string"
                }
            }
        },
        "dto.InstaResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Instagram"
                    }
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
	Host:        "127.0.0.1:4000",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Gostagram Swagger API",
	Description: "",
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
