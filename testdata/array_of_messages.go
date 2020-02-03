package testdata

const ArrayOfMessages = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "type": "string"
        },
        "payload": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "complete": {
                        "properties": {},
                        "type": "boolean"
                    },
                    "id": {
                        "properties": {},
                        "type": "integer"
                    },
                    "name": {
                        "properties": {},
                        "type": "string"
                    },
                    "rating": {
                        "properties": {},
                        "type": "number"
                    },
                    "timestamp": {
                        "properties": {},
                        "type": "string"
                    },
                    "topology": {
                        "enum": [
                            "FLAT",
                            0,
                            "NESTED_OBJECT",
                            1,
                            "NESTED_MESSAGE",
                            2,
                            "ARRAY_OF_TYPE",
                            3,
                            "ARRAY_OF_OBJECT",
                            4,
                            "ARRAY_OF_MESSAGE",
                            5
                        ],
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "additionalProperties": true,
                "type": "object"
            },
            "properties": {},
            "type": "array"
        }
    },
    "additionalProperties": true,
    "type": "object"
}`
