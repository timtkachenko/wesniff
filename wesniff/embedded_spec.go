// Code generated by go-swagger; DO NOT EDIT.

package wesniff

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
    "description": "api server",
    "title": "wesniff",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/events": {
      "post": {
        "security": [
          {
            "ApiKey": []
          }
        ],
        "tags": [
          "events"
        ],
        "summary": "accept callbacks",
        "operationId": "callbackHandle",
        "parameters": [
          {
            "description": "message and list of multiple queues",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EventBaseDto"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "type": "object",
                  "properties": {
                    "status": {
                      "type": "string"
                    }
                  }
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/health": {
      "x-swagger-router-controller": "health",
      "get": {
        "security": [],
        "description": "Healtcheck function",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [],
        "tags": [
          "users"
        ],
        "summary": "get Users",
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "description": "?limit=10,20",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/UserDto"
                  }
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-security-scopes": null
      },
      "post": {
        "security": [],
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "users"
        ],
        "summary": "create User",
        "operationId": "createUser",
        "parameters": [
          {
            "type": "file",
            "description": "Create User",
            "name": "image",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "email",
            "name": "email",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "$ref": "#/definitions/UserDto"
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AdditionalVerifiedData": {
      "type": "object",
      "properties": {
        "driversLicenseCategory": {
          "type": "object",
          "properties": {
            "B": {
              "type": "boolean"
            }
          }
        }
      }
    },
    "Document": {
      "type": "object",
      "properties": {
        "country": {
          "type": "string"
        },
        "number": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "validFrom": {
          "type": "string"
        },
        "validUntil": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "EventBase": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        },
        "verification": {
          "$ref": "#/definitions/Verification"
        }
      },
      "discriminator": "EventBase"
    },
    "EventBaseDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/EventBase"
        }
      ]
    },
    "Person": {
      "type": "object",
      "properties": {
        "citizenship": {
          "type": "string"
        },
        "dateOfBirth": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "gender": {
          "type": "string"
        },
        "idNumber": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "nationality": {
          "type": "string"
        },
        "pepSanctionMatch": {
          "type": "string"
        },
        "placeOfBirth": {
          "type": "string"
        },
        "yearOfBirth": {
          "type": "string"
        }
      }
    },
    "RiskLabel": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "label": {
          "type": "string"
        }
      }
    },
    "SuccessCode": {
      "description": "Status message for successful request",
      "type": "number",
      "example": 1
    },
    "UserBase": {
      "type": "object",
      "properties": {
        "role": {
          "description": "user role",
          "type": "string",
          "default": "admin",
          "enum": [
            "admin",
            "member"
          ]
        }
      },
      "discriminator": "UserBase"
    },
    "UserCreateDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        },
        {
          "required": [
            "email"
          ],
          "properties": {
            "email": {
              "description": "email",
              "type": "string",
              "format": "email",
              "x-go-custom-tag": "valid:\"email\"",
              "x-go-name": "email"
            }
          }
        }
      ]
    },
    "UserDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        },
        {
          "properties": {
            "email": {
              "description": "email",
              "type": "string",
              "x-omitempty": false
            },
            "id": {
              "description": "user id",
              "type": "string",
              "x-omitempty": false
            }
          }
        }
      ]
    },
    "UserUpdateDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        }
      ]
    },
    "Verification": {
      "type": "object",
      "properties": {
        "acceptanceTime": {
          "type": "string"
        },
        "additionalVerifiedData": {
          "$ref": "#/definitions/AdditionalVerifiedData"
        },
        "code": {
          "type": "integer"
        },
        "comments": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "decisionTime": {
          "type": "string"
        },
        "document": {
          "$ref": "#/definitions/Document"
        },
        "id": {
          "type": "string"
        },
        "person": {
          "$ref": "#/definitions/Person"
        },
        "reason": {
          "type": "string"
        },
        "reasonCode": {
          "type": "string"
        },
        "riskLabels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RiskLabel"
          }
        },
        "status": {
          "type": "string"
        },
        "technicalData": {
          "type": "object",
          "properties": {
            "ip": {
              "type": "string"
            }
          }
        },
        "vendorData": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "ApiKey": {
      "type": "apiKey",
      "name": "x-auth-client",
      "in": "header"
    },
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "Bearer": []
    }
  ]
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
    "description": "api server",
    "title": "wesniff",
    "version": "0.0.1"
  },
  "basePath": "/api",
  "paths": {
    "/events": {
      "post": {
        "security": [
          {
            "ApiKey": []
          }
        ],
        "tags": [
          "events"
        ],
        "summary": "accept callbacks",
        "operationId": "callbackHandle",
        "parameters": [
          {
            "description": "message and list of multiple queues",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EventBaseDto"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "type": "object",
                  "properties": {
                    "status": {
                      "type": "string"
                    }
                  }
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/health": {
      "x-swagger-router-controller": "health",
      "get": {
        "security": [],
        "description": "Healtcheck function",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "security": [],
        "tags": [
          "users"
        ],
        "summary": "get Users",
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "integer"
            },
            "description": "?limit=10,20",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/UserDto"
                  }
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        },
        "x-security-scopes": null
      },
      "post": {
        "security": [],
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "users"
        ],
        "summary": "create User",
        "operationId": "createUser",
        "parameters": [
          {
            "type": "file",
            "description": "Create User",
            "name": "image",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "email",
            "name": "email",
            "in": "formData",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "code": {
                  "$ref": "#/definitions/SuccessCode"
                },
                "data": {
                  "$ref": "#/definitions/UserDto"
                }
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AdditionalVerifiedData": {
      "type": "object",
      "properties": {
        "driversLicenseCategory": {
          "type": "object",
          "properties": {
            "B": {
              "type": "boolean"
            }
          }
        }
      }
    },
    "AdditionalVerifiedDataDriversLicenseCategory": {
      "type": "object",
      "properties": {
        "B": {
          "type": "boolean"
        }
      }
    },
    "CallbackHandleOKBodyData": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "Document": {
      "type": "object",
      "properties": {
        "country": {
          "type": "string"
        },
        "number": {
          "type": "string"
        },
        "type": {
          "type": "string"
        },
        "validFrom": {
          "type": "string"
        },
        "validUntil": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "EventBase": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        },
        "verification": {
          "$ref": "#/definitions/Verification"
        }
      },
      "discriminator": "EventBase"
    },
    "EventBaseDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/EventBase"
        }
      ]
    },
    "Person": {
      "type": "object",
      "properties": {
        "citizenship": {
          "type": "string"
        },
        "dateOfBirth": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "gender": {
          "type": "string"
        },
        "idNumber": {
          "type": "string"
        },
        "lastName": {
          "type": "string"
        },
        "nationality": {
          "type": "string"
        },
        "pepSanctionMatch": {
          "type": "string"
        },
        "placeOfBirth": {
          "type": "string"
        },
        "yearOfBirth": {
          "type": "string"
        }
      }
    },
    "RiskLabel": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "label": {
          "type": "string"
        }
      }
    },
    "SuccessCode": {
      "description": "Status message for successful request",
      "type": "number",
      "example": 1
    },
    "UserBase": {
      "type": "object",
      "properties": {
        "role": {
          "description": "user role",
          "type": "string",
          "default": "admin",
          "enum": [
            "admin",
            "member"
          ]
        }
      },
      "discriminator": "UserBase"
    },
    "UserCreateDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        },
        {
          "required": [
            "email"
          ],
          "properties": {
            "email": {
              "description": "email",
              "type": "string",
              "format": "email",
              "x-go-custom-tag": "valid:\"email\"",
              "x-go-name": "email"
            }
          }
        }
      ]
    },
    "UserDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        },
        {
          "properties": {
            "email": {
              "description": "email",
              "type": "string",
              "x-omitempty": false
            },
            "id": {
              "description": "user id",
              "type": "string",
              "x-omitempty": false
            }
          }
        }
      ]
    },
    "UserUpdateDto": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/UserBase"
        }
      ]
    },
    "Verification": {
      "type": "object",
      "properties": {
        "acceptanceTime": {
          "type": "string"
        },
        "additionalVerifiedData": {
          "$ref": "#/definitions/AdditionalVerifiedData"
        },
        "code": {
          "type": "integer"
        },
        "comments": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "decisionTime": {
          "type": "string"
        },
        "document": {
          "$ref": "#/definitions/Document"
        },
        "id": {
          "type": "string"
        },
        "person": {
          "$ref": "#/definitions/Person"
        },
        "reason": {
          "type": "string"
        },
        "reasonCode": {
          "type": "string"
        },
        "riskLabels": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/RiskLabel"
          }
        },
        "status": {
          "type": "string"
        },
        "technicalData": {
          "type": "object",
          "properties": {
            "ip": {
              "type": "string"
            }
          }
        },
        "vendorData": {
          "type": "string"
        }
      }
    },
    "VerificationTechnicalData": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "ApiKey": {
      "type": "apiKey",
      "name": "x-auth-client",
      "in": "header"
    },
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "Bearer": []
    }
  ]
}`))
}