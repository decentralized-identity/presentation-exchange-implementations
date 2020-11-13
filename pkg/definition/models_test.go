package definition

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decentralized-identity/presentation-exchange-implementations/internal/schema"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

func TestPresentationDefinition(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		samplePresDef, err := testcases.GetJSONFile(testcases.BasicPresentationDefinition)
		assert.NoError(t, err)

		var presDef PresentationDefinitionHolder
		presDefBytes := []byte(samplePresDef)
		err = json.Unmarshal(presDefBytes, &presDef)
		assert.NoError(t, err)

		assert.NoError(t, util.Validate(presDef))

		// Roundtrip and compare
		roundTripBytes, err := json.Marshal(presDef)
		assert.NoError(t, err)

		true, err := schema.JSONBytesEqual(presDefBytes, roundTripBytes)
		assert.NoError(t, err)
		assert.True(t, true)
	})

	t.Run("Single Group", func(t *testing.T) {
		samplePresDef, err := testcases.GetJSONFile(testcases.SingleGroupPresentationDefinition)
		assert.NoError(t, err)

		var presDef PresentationDefinitionHolder
		presDefBytes := []byte(samplePresDef)
		err = json.Unmarshal(presDefBytes, &presDef)
		assert.NoError(t, err)

		assert.NoError(t, util.Validate(presDef))

		// Roundtrip and compare
		roundTripBytes, err := json.Marshal(presDef)
		assert.NoError(t, err)

		true, err := schema.JSONBytesEqual(presDefBytes, roundTripBytes)
		assert.NoError(t, err)
		assert.True(t, true)
	})

	t.Run("Multi Group", func(t *testing.T) {
		samplePresDef, err := testcases.GetJSONFile(testcases.MultiGroupPresentationDefinition)
		assert.NoError(t, err)

		var presDef PresentationDefinitionHolder
		presDefBytes := []byte(samplePresDef)
		err = json.Unmarshal(presDefBytes, &presDef)
		assert.NoError(t, err)

		assert.NoError(t, util.Validate(presDef))

		// Roundtrip and compare
		roundTripBytes, err := json.Marshal(presDef)
		assert.NoError(t, err)

		true, err := schema.JSONBytesEqual(presDefBytes, roundTripBytes)
		assert.NoError(t, err)
		assert.True(t, true)
	})
}

func TestPresentationDefinitionBuilder(t *testing.T) {
	t.Run("one ldp", func(t *testing.T) {
		// minimally complaint pres def
		b := NewPresentationDefinitionBuilder()
		err := b.AddInputDescriptor(InputDescriptor{
			ID:          "test",
			Schema:      &Schema{
				URI:      []string{"test"},
				Name:     "test",
			},
		})
		assert.NoError(t, err)

		err = b.SetLDPFormat(LDP, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		pres, err := b.Build()
		assert.NoError(t, err)
		assert.NotNil(t, pres)
	})

	t.Run("one jwt", func(t *testing.T) {
		// minimally complaint pres def
		b := NewPresentationDefinitionBuilder()
		err := b.AddInputDescriptor(InputDescriptor{
			ID:          "test",
			Schema:      &Schema{
				URI:      []string{"test"},
				Name:     "test",
			},
		})
		assert.NoError(t, err)

		err = b.SetJWTFormat(JWT, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		pres, err := b.Build()
		assert.NoError(t, err)
		assert.NotNil(t, pres)
	})

	t.Run("mixed ldp and jwt", func(t *testing.T) {
		// minimally complaint pres def
		b := NewPresentationDefinitionBuilder()
		err := b.AddInputDescriptor(InputDescriptor{
			ID:          "test",
			Schema:      &Schema{
				URI:      []string{"test"},
				Name:     "test",
			},
		})
		assert.NoError(t, err)

		err = b.SetJWTFormat(JWT, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)

		err = b.SetLDPFormat(LDP, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)

		pres, err := b.Build()
		assert.NoError(t, err)
		assert.NotNil(t, pres)
	})

	t.Run("mixed ldps and jwts", func(t *testing.T) {
		// minimally complaint pres def
		b := NewPresentationDefinitionBuilder()
		err := b.AddInputDescriptor(InputDescriptor{
			ID:          "test",
			Schema:      &Schema{
				URI:      []string{"test"},
				Name:     "test",
			},
		})
		assert.NoError(t, err)

		err = b.SetJWTFormat(JWT, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		err = b.SetJWTFormat(JWTVC, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		err = b.SetJWTFormat(JWTVP, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)

		err = b.SetLDPFormat(LDP, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		err = b.SetLDPFormat(LDPVC, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)
		err = b.SetLDPFormat(LDPVP, []string{"Ed25519Signature2018"})
		assert.NoError(t, err)

		pres, err := b.Build()
		assert.NoError(t, err)
		assert.NotNil(t, pres)
	})
}