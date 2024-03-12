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
        "/buku": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all buku",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buku"
                ],
                "summary": "Get all buku",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controller.Buku"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create buku",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buku"
                ],
                "summary": "Create buku",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Cover buku",
                        "name": "cover",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID Kategori",
                        "name": "id_kategori",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID Rak",
                        "name": "id_rak",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Judul buku",
                        "name": "judul",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Penulis buku",
                        "name": "penulis",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Penerbit buku",
                        "name": "penerbit",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tahun terbit buku",
                        "name": "tahun_terbit",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Jumlah halaman buku",
                        "name": "jml_hal",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Deskripsi buku",
                        "name": "deskripsi",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.Buku"
                        }
                    }
                }
            }
        },
        "/buku/{id}": {
            "get": {
                "description": "Get buku by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buku"
                ],
                "summary": "Get buku by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Buku ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Buku"
                        }
                    }
                }
            },
            "put": {
                "description": "Update buku",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buku"
                ],
                "summary": "Update buku",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Buku ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Buku",
                        "name": "buku",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Buku"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Buku"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete buku",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buku"
                ],
                "summary": "Delete buku",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Buku ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Buku"
                        }
                    }
                }
            }
        },
        "/kategori": {
            "get": {
                "description": "Get all kategori buku",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Get all kategori",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controller.Kategori"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create kategori",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Create kategori",
                "parameters": [
                    {
                        "description": "Kategori",
                        "name": "kategori",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                }
            }
        },
        "/kategori/{id}": {
            "get": {
                "description": "Get kategori by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Get kategori by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Kategori ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                }
            },
            "put": {
                "description": "Update kategori",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Update kategori",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Kategori ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Kategori",
                        "name": "kategori",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete kategori",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kategori"
                ],
                "summary": "Delete kategori",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Kategori ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Kategori"
                        }
                    }
                }
            }
        },
        "/rak": {
            "post": {
                "description": "Create rak",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rak"
                ],
                "summary": "Create rak",
                "parameters": [
                    {
                        "description": "Rak",
                        "name": "rak",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                }
            }
        },
        "/rak/{id}": {
            "get": {
                "description": "Get rak by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rak"
                ],
                "summary": "Get rak by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rak ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                }
            },
            "put": {
                "description": "Update rak",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rak"
                ],
                "summary": "Update rak",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rak ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Rak",
                        "name": "rak",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete rak",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rak"
                ],
                "summary": "Delete rak",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rak ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Rak"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Buku": {
            "type": "object",
            "properties": {
                "cover": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "id_kategori": {
                    "type": "integer"
                },
                "id_rak": {
                    "type": "integer"
                },
                "jml_hal": {
                    "type": "integer"
                },
                "judul": {
                    "type": "string"
                },
                "nama_kategori": {
                    "type": "string"
                },
                "nama_rak": {
                    "type": "string"
                },
                "penerbit": {
                    "type": "string"
                },
                "penulis": {
                    "type": "string"
                },
                "tahun_terbit": {
                    "type": "string"
                }
            }
        },
        "controller.Kategori": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nama_kategori": {
                    "type": "string"
                }
            }
        },
        "controller.Rak": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nama_rak": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.6",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "PerpusDigital",
	Description:      "Backend Perpustakaan Digital",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
