// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "XuanWuService",
    "version": "1.0.0"
  },
  "basePath": "/xuanwu/v1",
  "paths": {
    "/hello": {
      "get": {
        "description": "Hello",
        "summary": "Hello",
        "operationId": "hello",
        "responses": {
          "200": {
            "description": "Hello",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "BadRequest",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/hello/trace": {
      "get": {
        "description": "Hello",
        "summary": "HelloTrace",
        "operationId": "helloTrace",
        "responses": {
          "200": {
            "description": "Hello",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "BadRequest",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Meta": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "integer"
        },
        "updated_at": {
          "type": "integer"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "required": [
        "limit",
        "offset",
        "total"
      ],
      "properties": {
        "limit": {
          "type": "integer"
        },
        "offset": {
          "type": "integer"
        },
        "total": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "XuanWuService",
    "version": "1.0.0"
  },
  "basePath": "/xuanwu/v1",
  "paths": {
    "/hello": {
      "get": {
        "description": "Hello",
        "summary": "Hello",
        "operationId": "hello",
        "responses": {
          "200": {
            "description": "Hello",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "BadRequest",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/hello/trace": {
      "get": {
        "description": "Hello",
        "summary": "HelloTrace",
        "operationId": "helloTrace",
        "responses": {
          "200": {
            "description": "Hello",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "BadRequest",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Meta": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "integer"
        },
        "updated_at": {
          "type": "integer"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "required": [
        "limit",
        "offset",
        "total"
      ],
      "properties": {
        "limit": {
          "type": "integer"
        },
        "offset": {
          "type": "integer"
        },
        "total": {
          "type": "integer"
        }
      }
    }
  }
}`))
}