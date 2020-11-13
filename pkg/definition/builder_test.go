package definition

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

// https://identity.foundation/presentation-exchange/#presentation-definition---basic-example
func TestPresentationDefinitionBuilder_BasicExample(t *testing.T) {
	b := NewPresentationDefinitionBuilder()

	// shouldn't validate as empty
	_, err := b.Build()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation for 'InputDescriptors' failed on the 'required' tag")

	// create an input descriptor
	id1 := NewInputDescriptor("banking_input")
	err = id1.SetSchema(Schema{
		URI:     []string{"https://bank-standards.com/customer.json"},
		Name:    "Bank Account Information",
		Purpose: "We need your bank and account information.",
	})
	assert.NoError(t, err)

	field := NewConstraintsField([]string{"$.issuer", "$.vc.issuer", "$.iss"})
	field.SetPurpose("The credential must be from one of the specified issuers")
	err = field.SetFilter(Filter{
		Type:    "string",
		Pattern: "did:example:123|did:example:456",
	})
	assert.NoError(t, err)

	// now build
	err = id1.SetConstraints(*field)
	assert.NoError(t, err)

	id1.SetConstraintsLimitDisclosure(true)

	// validate
	err = util.Validate(id1)
	assert.NoError(t, err)

	// add the descriptor to the builder
	err = b.AddInputDescriptor(*id1)
	assert.NoError(t, err)

	// add a second input descriptor
	id2 := NewInputDescriptor("citizenship_input")
	err = id2.SetSchema(Schema{
		URI:  []string{"hub://did:foo:123/Collections/schema.us.gov/passport.json"},
		Name: "US Passport",
	})
	assert.NoError(t, err)

	field2 := NewConstraintsField([]string{"$.credentialSubject.birth_date", "$.vc.credentialSubject.birth_date", "$.birth_date"})
	err = field2.SetFilter(Filter{
		Type:    "date",
		Minimum: "1999-5-16",
	})
	assert.NoError(t, err)

	err = id2.SetConstraints(*field2)
	assert.NoError(t, err)

	// validate
	err = util.Validate(id2)
	assert.NoError(t, err)

	// add the descriptor to the builder
	err = b.AddInputDescriptor(*id2)
	assert.NoError(t, err)

	presDef, err := b.Build()
	assert.NoError(t, err)

	assert.NoError(t, util.Validate(presDef))

	presDefJSON, err := util.ToJSON(presDef)
	assert.NoError(t, err)

	// get sample json from packr
	testPresDefJSON, err := testcases.GetJSONFile(testcases.BasicPresentationDefinition)
	assert.NoError(t, err)

	// Make sure our builder has the same result
	same, err := util.CompareJSON(presDefJSON, testPresDefJSON)
	assert.NoError(t, err)
	assert.True(t, same)
}

// TODO as the spec is in flux the remaining tests will be implemented once v0.1.0 is finalized

// // https://identity.foundation/presentation-exchange/#presentation-definition---single-group-example
// func TestPresentationDefinitionBuilder_SingleGroupExample(t *testing.T) {
// }
//
// // https://identity.foundation/presentation-exchange/#presentation-definition---multi-group-example
// func TestPresentationDefinitionBuilder_MultiGroupExample(t *testing.T) {
// }
