// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/acc-data/get/{tag}": {
            "get": {
                "description": "gets the latest data of the specified tag - get your tags via \"list\" first",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acc-data"
                ],
                "summary": "gets latest data of the specified tag",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"12:34:56:78:90:12\"",
                        "description": "the tags address",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AccelerationSensor"
                            }
                        }
                    }
                }
            }
        },
        "/acc-data/get/{tag}/{page}": {
            "get": {
                "description": "gets data of the specified tag - get your tags via \"list\" first",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "acc-data"
                ],
                "summary": "gets acc data of the specified tag",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"12:34:56:78:90:12\"",
                        "description": "the tags address",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "the page of measurements you would like to get",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/config/get/{gatewayid}": {
            "get": {
                "description": "get config of a specific gateway to be able to decide on version updates etc. To find the gateways id use /structure/gateway/list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "lists available configs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the gateways id",
                        "name": "gatewayid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.GatewayConfig"
                            }
                        }
                    }
                }
            }
        },
        "/config/get/{gatewayid}/{tagaddress}": {
            "get": {
                "description": "get config of a specific gateway to be able to decide on version updates etc. To find the gateways id use /structure/gateway/list, to get the tags address use /structure/tag/list/{gatewayid}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "lists available configs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the gateways id",
                        "name": "gatewayid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the tags id",
                        "name": "tagaddress",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.GatewayConfig"
                            }
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/structure/gateway/list": {
            "get": {
                "description": "lists gateways",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "structure"
                ],
                "summary": "lists available gateways",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Gateway"
                            }
                        }
                    }
                }
            }
        },
        "/structure/tag/list/{gatewayid}": {
            "get": {
                "description": "lists tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "structure"
                ],
                "summary": "lists available tags - get the gateway_id from structure/gateway/list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the gateway",
                        "name": "gatewayid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Gateway"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AccelerationSensor": {
            "type": "object",
            "properties": {
                "measurements": {
                    "type": "array",
                    "items": {}
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Gateway": {
            "description": "tags - all the tags, known to the gateway config - GatewayConfig that belongs to this Gateway network_segment - subnet the Gateway is in last_contact - last time the api heard back from the Gateway online - bool that determines if the Gateway is currently online ip_address - current IPv4 of the Gateway id - unique identifier",
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/model.GatewayConfig"
                },
                "id": {
                    "type": "string"
                },
                "ip_address": {
                    "type": "string"
                },
                "last_contact": {
                    "type": "string"
                },
                "network_segment": {
                    "type": "integer"
                },
                "online": {
                    "type": "boolean"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                }
            }
        },
        "model.GatewayConfig": {
            "description": "version - version of the gateway's firmware poll_interval - how often the tags will be polled via bluetooth max_allowed_clients - how many tags will be accepted api_timeout - how long it takes for the gateway to restart when the api-server is unavailable",
            "type": "object",
            "properties": {
                "api_timeout": {
                    "type": "number"
                },
                "max_allowed_clients": {
                    "type": "integer"
                },
                "poll_interval": {
                    "type": "integer"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.Tag": {
            "description": "sensors - all the sensors, mounted on this tag (list) address - unique MAC-address of the tag (bluetooth MAC) name - name of the tag that derives from the MAC-address last_contact - last time the Gateway heard back from the Tag online - bool that determines if the Tag is currently online config - TagConfig that belongs to this Tag",
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "config": {
                    "$ref": "#/definitions/model.TagConfig"
                },
                "last_contact": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "online": {
                    "type": "boolean"
                },
                "sensors": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "model.TagConfig": {
            "description": "samplerate - samplerate of the tag scan_interval - interval the sensors of the tag will be polled (ms) resolution - bit depth resolution of the sensors scale - scaling factor for values from the sensors (for compression) dsp_function - dsp function for signal evaluation (enum) dsp_parameter - dsp configuration parameter (enum) mode - current measurement mode (enum) divider - divider for the samplerate",
            "type": "object",
            "properties": {
                "divider": {
                    "type": "integer"
                },
                "dsp_function": {
                    "type": "integer"
                },
                "dsp_parameter": {
                    "type": "integer"
                },
                "mode": {
                    "type": "integer"
                },
                "resolution": {
                    "type": "integer"
                },
                "samplerate": {
                    "type": "integer"
                },
                "scale": {
                    "type": "integer"
                },
                "scan_interval": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
