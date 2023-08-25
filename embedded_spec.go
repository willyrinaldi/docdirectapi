// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This file is auto-generated using go-swagger annotations.",
    "title": "Rasio Keuangan API",
    "contact": {
      "name": "Your Name",
      "email": "your@email.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8082",
  "basePath": "/",
  "paths": {
    "/direct/v1/saham/dp/eq/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Data History"
        ],
        "summary": "Get Data History",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "startDate",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "endDate",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object",
              "properties": {
                "bottom": {
                  "type": "object",
                  "properties": {
                    "Average": {
                      "type": "number"
                    },
                    "Change": {
                      "type": "number"
                    },
                    "Difference": {
                      "type": "integer"
                    },
                    "Highest": {
                      "type": "integer"
                    },
                    "LastUpdate": {
                      "type": "string"
                    },
                    "Lowest": {
                      "type": "integer"
                    }
                  }
                },
                "changes": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "Change": {
                        "type": "number"
                      },
                      "Close": {
                        "type": "integer"
                      },
                      "Date": {
                        "type": "string"
                      },
                      "Freq": {
                        "type": "integer"
                      },
                      "High": {
                        "type": "integer"
                      },
                      "Low": {
                        "type": "integer"
                      },
                      "Open": {
                        "type": "integer"
                      },
                      "Val": {
                        "type": "integer"
                      },
                      "Vol": {
                        "type": "integer"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "application/json": {
                "average": 6725.262129032257,
                "change": 3.731814563003012,
                "data": [
                  {
                    "change": 0.24202422165240495,
                    "close": 6880.802,
                    "date": "2023-07-21",
                    "freq": 1060985,
                    "high": 6880.802,
                    "low": 6837.646,
                    "prev": 6864.189,
                    "value": 8617503732103,
                    "volume": 16262882812
                  }
                ],
                "difference": 352.5159999999996,
                "highest": 6931.272,
                "lastUpdate": "2023-07-21",
                "lowest": 6578.756
              }
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "405": {
            "description": "Method Not Allowed"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/dp/ix/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Index Data"
        ],
        "summary": "Get Index Data",
        "parameters": [
          {
            "type": "string",
            "name": "indexCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "startDate",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "endDate",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/fdr/ip/get-informasi-pasar-data": {
      "get": {
        "description": "Get Informasi Pasar Data based on provided parameters.",
        "tags": [
          "Informasi Pasar"
        ],
        "summary": "Get Informasi Pasar Data",
        "parameters": [
          {
            "type": "string",
            "description": "Emiten code",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Year of the data",
            "name": "periode",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Quarter of the data",
            "name": "q",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Granularity (annually/quarterly)",
            "name": "granularity",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResultIndexDataIinfPasar"
              }
            }
          },
          "400": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      }
    },
    "/direct/v1/saham/fdr/rk/get-rasio-keuangan-data": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Rasio Keuangan"
        ],
        "summary": "Get Rasio Keuangan Data",
        "parameters": [
          {
            "type": "string",
            "description": "Emiten code",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Year of the data",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Quarter of the data",
            "name": "q",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Granularity (annually/quarterly)",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResultIndexData"
              }
            }
          },
          "400": {
            "$ref": "#/responses/ErrorResponse"
          }
        }
      }
    },
    "/direct/v1/saham/lk/bs/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Balance Sheet Data"
        ],
        "summary": "Get Balance Sheet Data",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "q",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/lk/is/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Income Statement"
        ],
        "summary": "Get Income Statement Data",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "q",
            "in": "query"
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "ResultIndexData": {
      "type": "object",
      "properties": {
        "ATO": {
          "type": "number",
          "format": "double"
        },
        "BVPS": {
          "type": "number",
          "format": "double"
        },
        "DAR": {
          "type": "number",
          "format": "double"
        },
        "DER": {
          "type": "number",
          "format": "double"
        },
        "EVPerEbitTTM": {
          "type": "number",
          "format": "double"
        },
        "EVPerEbitdaTTM": {
          "type": "number",
          "format": "double"
        },
        "EVPerRevenueTTM": {
          "type": "number",
          "format": "double"
        },
        "GPM": {
          "type": "number",
          "format": "double"
        },
        "NPM": {
          "type": "number",
          "format": "double"
        },
        "OPM": {
          "type": "number",
          "format": "double"
        },
        "PBV": {
          "type": "number",
          "format": "double"
        },
        "cashRatio": {
          "type": "number",
          "format": "double"
        },
        "currentRatio": {
          "type": "number",
          "format": "double"
        },
        "eps": {
          "type": "number",
          "format": "double"
        },
        "pe": {
          "type": "number",
          "format": "double"
        },
        "period": {
          "type": "string"
        },
        "ps": {
          "type": "number",
          "format": "double"
        },
        "revenuePerShares": {
          "type": "number",
          "format": "double"
        },
        "roa": {
          "type": "number",
          "format": "double"
        },
        "roe": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "ResultIndexDataIinfPasar": {
      "type": "object",
      "properties": {
        "close": {
          "type": "integer",
          "format": "int64"
        },
        "enterpriseValue": {
          "type": "number",
          "format": "double"
        },
        "listedShares": {
          "type": "integer",
          "format": "int64"
        },
        "marketCap": {
          "type": "integer",
          "format": "int64"
        },
        "peringkatKapitalisasiPasar": {
          "type": "integer",
          "format": "int64"
        },
        "period": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "ErrorResponse": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "APIs for retrieving Rasio Keuangan data",
      "name": "Rasio Keuangan"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This file is auto-generated using go-swagger annotations.",
    "title": "Rasio Keuangan API",
    "contact": {
      "name": "Your Name",
      "email": "your@email.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8082",
  "basePath": "/",
  "paths": {
    "/direct/v1/saham/dp/eq/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Data History"
        ],
        "summary": "Get Data History",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "startDate",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "endDate",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object",
              "properties": {
                "bottom": {
                  "type": "object",
                  "properties": {
                    "Average": {
                      "type": "number"
                    },
                    "Change": {
                      "type": "number"
                    },
                    "Difference": {
                      "type": "integer"
                    },
                    "Highest": {
                      "type": "integer"
                    },
                    "LastUpdate": {
                      "type": "string"
                    },
                    "Lowest": {
                      "type": "integer"
                    }
                  }
                },
                "changes": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/ChangesItems0"
                  }
                }
              }
            },
            "examples": {
              "application/json": {
                "average": 6725.262129032257,
                "change": 3.731814563003012,
                "data": [
                  {
                    "change": 0.24202422165240495,
                    "close": 6880.802,
                    "date": "2023-07-21",
                    "freq": 1060985,
                    "high": 6880.802,
                    "low": 6837.646,
                    "prev": 6864.189,
                    "value": 8617503732103,
                    "volume": 16262882812
                  }
                ],
                "difference": 352.5159999999996,
                "highest": 6931.272,
                "lastUpdate": "2023-07-21",
                "lowest": 6578.756
              }
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "405": {
            "description": "Method Not Allowed"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/dp/ix/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Index Data"
        ],
        "summary": "Get Index Data",
        "parameters": [
          {
            "type": "string",
            "name": "indexCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "startDate",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "name": "endDate",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/fdr/ip/get-informasi-pasar-data": {
      "get": {
        "description": "Get Informasi Pasar Data based on provided parameters.",
        "tags": [
          "Informasi Pasar"
        ],
        "summary": "Get Informasi Pasar Data",
        "parameters": [
          {
            "type": "string",
            "description": "Emiten code",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Year of the data",
            "name": "periode",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Quarter of the data",
            "name": "q",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Granularity (annually/quarterly)",
            "name": "granularity",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResultIndexDataIinfPasar"
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/direct/v1/saham/fdr/rk/get-rasio-keuangan-data": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Rasio Keuangan"
        ],
        "summary": "Get Rasio Keuangan Data",
        "parameters": [
          {
            "type": "string",
            "description": "Emiten code",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Year of the data",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Quarter of the data",
            "name": "q",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Granularity (annually/quarterly)",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResultIndexData"
              }
            }
          },
          "400": {
            "description": "Error response",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/direct/v1/saham/lk/bs/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Balance Sheet Data"
        ],
        "summary": "Get Balance Sheet Data",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "q",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Not Found"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/direct/v1/saham/lk/is/": {
      "get": {
        "description": "Get Rasio Keuangan Data based on provided parameters.",
        "tags": [
          "Income Statement"
        ],
        "summary": "Get Income Statement Data",
        "parameters": [
          {
            "type": "string",
            "name": "secCode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "periode",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "q",
            "in": "query"
          },
          {
            "type": "string",
            "name": "granularity",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "schema": {
              "type": "array",
              "items": {
                "type": "object"
              }
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "ChangesItems0": {
      "type": "object",
      "properties": {
        "Change": {
          "type": "number"
        },
        "Close": {
          "type": "integer"
        },
        "Date": {
          "type": "string"
        },
        "Freq": {
          "type": "integer"
        },
        "High": {
          "type": "integer"
        },
        "Low": {
          "type": "integer"
        },
        "Open": {
          "type": "integer"
        },
        "Val": {
          "type": "integer"
        },
        "Vol": {
          "type": "integer"
        }
      }
    },
    "GetDirectV1SahamDpEqOKBodyBottom": {
      "type": "object",
      "properties": {
        "Average": {
          "type": "number"
        },
        "Change": {
          "type": "number"
        },
        "Difference": {
          "type": "integer"
        },
        "Highest": {
          "type": "integer"
        },
        "LastUpdate": {
          "type": "string"
        },
        "Lowest": {
          "type": "integer"
        }
      }
    },
    "ResultIndexData": {
      "type": "object",
      "properties": {
        "ATO": {
          "type": "number",
          "format": "double"
        },
        "BVPS": {
          "type": "number",
          "format": "double"
        },
        "DAR": {
          "type": "number",
          "format": "double"
        },
        "DER": {
          "type": "number",
          "format": "double"
        },
        "EVPerEbitTTM": {
          "type": "number",
          "format": "double"
        },
        "EVPerEbitdaTTM": {
          "type": "number",
          "format": "double"
        },
        "EVPerRevenueTTM": {
          "type": "number",
          "format": "double"
        },
        "GPM": {
          "type": "number",
          "format": "double"
        },
        "NPM": {
          "type": "number",
          "format": "double"
        },
        "OPM": {
          "type": "number",
          "format": "double"
        },
        "PBV": {
          "type": "number",
          "format": "double"
        },
        "cashRatio": {
          "type": "number",
          "format": "double"
        },
        "currentRatio": {
          "type": "number",
          "format": "double"
        },
        "eps": {
          "type": "number",
          "format": "double"
        },
        "pe": {
          "type": "number",
          "format": "double"
        },
        "period": {
          "type": "string"
        },
        "ps": {
          "type": "number",
          "format": "double"
        },
        "revenuePerShares": {
          "type": "number",
          "format": "double"
        },
        "roa": {
          "type": "number",
          "format": "double"
        },
        "roe": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "ResultIndexDataIinfPasar": {
      "type": "object",
      "properties": {
        "close": {
          "type": "integer",
          "format": "int64"
        },
        "enterpriseValue": {
          "type": "number",
          "format": "double"
        },
        "listedShares": {
          "type": "integer",
          "format": "int64"
        },
        "marketCap": {
          "type": "integer",
          "format": "int64"
        },
        "peringkatKapitalisasiPasar": {
          "type": "integer",
          "format": "int64"
        },
        "period": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "ErrorResponse": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    }
  },
  "tags": [
    {
      "description": "APIs for retrieving Rasio Keuangan data",
      "name": "Rasio Keuangan"
    }
  ]
}`))
}