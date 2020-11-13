package schema

import (
	"encoding/json"
	"reflect"

	"github.com/sirupsen/logrus"
)

// go representation of json schema document
type JSONSchemaMap map[string]interface{}

type Properties map[string]interface{}

// Assumes the json schema has a properties field
func (j JSONSchemaMap) Properties() Properties {
	if properties, ok := j["properties"]; ok {
		return properties.(map[string]interface{})
	}
	return map[string]interface{}{}
}

// Assumes the json schema has a description field
func (j JSONSchemaMap) Description() string {
	if description, ok := j["description"]; ok {
		return description.(string)
	}
	return ""
}

func (j JSONSchemaMap) AllowsAdditionalProperties() bool {
	if v, exists := j["additionalProperties"]; exists {
		if additionalProps, ok := v.(bool); ok {
			return additionalProps
		}
	}
	return false
}

func (j JSONSchemaMap) RequiredFields() []string {
	if v, exists := j["required"]; exists {
		if requiredFields, ok := v.([]interface{}); ok {
			required := make([]string, 0, len(requiredFields))
			for _, f := range requiredFields {
				required = append(required, f.(string))
			}
			return required
		}
	}
	return []string{}
}

func Type(field interface{}) string {
	if asMap, isMap := field.(map[string]interface{}); isMap {
		if v, exists := asMap["type"]; exists {
			if typeString, ok := v.(string); ok {
				return typeString
			}
		}
	}
	return ""
}

func Format(field interface{}) string {
	if asMap, isMap := field.(map[string]interface{}); isMap {
		if v, exists := asMap["format"]; exists {
			if formatString, ok := v.(string); ok {
				return formatString
			}
		}
	}
	return ""
}

func Contains(field string, required []string) bool {
	for _, f := range required {
		if f == field {
			return true
		}
	}
	return false
}

func (j JSONSchemaMap) ToJSON() string {
	bytes, err := json.Marshal(j)
	if err != nil {
		logrus.WithError(err).Error("Unable to jsonify schema")
		panic(err)
	}
	return string(bytes)
}

// JSONBytesEqual compares the JSON in two byte slices for deep equality, ignoring whitespace
// and other non-semantically meaningful formatting differences.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j1, j2 interface{}
	if err := json.Unmarshal(a, &j1); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j1), nil
}
