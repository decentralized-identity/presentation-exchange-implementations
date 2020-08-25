package schema

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

// Validate exists to hide gojsonschema logic within this file
// it is the entry-point to validation logic, requiring the caller pass in valid json strings for each argument
func Validate(schema, document string) error {
	if !isJSON(schema) {
		return fmt.Errorf("schema is not valid json: %s", schema)
	} else if !isJSON(document) {
		return fmt.Errorf("document is not valid json: %s", document)
	}
	return validateWithJSONLoader(gojsonschema.NewStringLoader(schema), gojsonschema.NewStringLoader(document))
}

// validateWithJSONLoader takes schema and document loaders; the document from the loader is validated against
// the schema from the loader. Nil if good, error if bad
func validateWithJSONLoader(schemaLoader, documentLoader gojsonschema.JSONLoader) error {
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		logrus.WithError(err).Error("failed to validateWithJSONLoader document against s")
		return err
	}

	if !result.Valid() {
		// Accumulate errs
		var errs []string
		for _, err := range result.Errors() {
			errs = append(errs, err.String())
		}
		return fmt.Errorf("schema failed validation: %s", strings.Join(errs, ","))
	}
	return nil
}

// ValidateJSONSchema takes in a string that is purported to be a JSON schema (schema definition)
// An error is returned if it is not a valid JSON schema, and nil is returned on success
func ValidateJSONSchema(maybeSchema JSONSchemaMap) error {
	schemaLoader := gojsonschema.NewSchemaLoader()
	schemaLoader.Validate = true
	return schemaLoader.AddSchemas(gojsonschema.NewStringLoader(maybeSchema.ToJSON()))
}

// True if string is valid JSON, false otherwise
func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
