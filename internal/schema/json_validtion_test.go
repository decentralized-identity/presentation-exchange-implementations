package schema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	definition "github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition/testcases"
	submission "github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission/verifiablepresentation/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/schemas"
)

func TestPresentationDefinitionSchema(t *testing.T) {
	presDefSchema, err := schemas.GetJSONFile(schemas.PresentationDefinition)
	assert.NoError(t, err)

	t.Run("Validate the schema is valid", func(t *testing.T) {
		var schemaMap JSONSchemaMap
		err := json.Unmarshal([]byte(presDefSchema), &schemaMap)
		assert.NoError(t, err)

		err = ValidateJSONSchema(schemaMap)
		assert.NoError(t, err)
	})

	t.Run("Validate a sample definition is valid against the schema", func(t *testing.T) {
		samplePresDef, err := definition.GetJSONFile(definition.BasicPresentationDefinition)
		assert.NoError(t, err)

		// Validate against presDefSchema
		assert.NoError(t, Validate(presDefSchema, samplePresDef))
	})
}

func TestPresentationSubmissionSchema(t *testing.T) {
	presSubSchema, err := schemas.GetJSONFile(schemas.PresentationSubmission)
	assert.NoError(t, err)

	t.Run("Validate the schema is valid", func(t *testing.T) {
		var schemaMap JSONSchemaMap
		err := json.Unmarshal([]byte(presSubSchema), &schemaMap)
		assert.NoError(t, err)

		err = ValidateJSONSchema(schemaMap)
		assert.NoError(t, err)
	})

	t.Run("Validate a sample submission is valid against the schema", func(t *testing.T) {
		samplePresSub, err := submission.GetJSONFile(submission.SampleSubmission)
		assert.NoError(t, err)

		// Validate against presSubSchema
		assert.NoError(t, Validate(presSubSchema, samplePresSub))
	})
}
