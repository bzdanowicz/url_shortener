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
        "/url": {
            "post": {
                "description": "shorten original link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creating new short link.",
                "parameters": [
                    {
                        "description": "OriginalUrl",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.UrlCreationRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Returns shorten url",
                        "schema": {
                            "$ref": "#/definitions/main.UrlCreationResponse"
                        }
                    },
                    "400": {
                        "description": "Failure response",
                        "schema": {
                            "$ref": "#/definitions/main.UrlCreationResponse"
                        }
                    }
                }
            }
        },
        "/{url}": {
            "get": {
                "description": "redirect to the original page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Link redirection.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "url",
                        "name": "url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "redirect to original page"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.UrlCreationRequest": {
            "type": "object",
            "required": [
                "original_url"
            ],
            "properties": {
                "original_url": {
                    "type": "string"
                }
            }
        },
        "main.UrlCreationResponse": {
            "type": "object",
            "required": [
                "message",
                "new_url",
                "original_url"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "new_url": {
                    "type": "string"
                },
                "original_url": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
	swag.Register(swag.Name, &s{})
}
