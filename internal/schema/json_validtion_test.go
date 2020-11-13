package schema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	definition "go.wday.io/credentials-open-source/presentation-exchange/pkg/definition/testcases"
	submission "go.wday.io/credentials-open-source/presentation-exchange/pkg/submission/verifiablepresentation/testcases"
	"go.wday.io/credentials-open-source/presentation-exchange/schemas"
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
