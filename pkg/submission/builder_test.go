package submission

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission/verifiablepresentation/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

// https://identity.foundation/presentation-exchange/#presentation-submission---verifiable-presentation
func TestPresentationSubmissionBuilder(t *testing.T) {
	b := NewPresentationSubmissionBuilder("32f54163-7166-48f1-93d8-ff217bdb0653")
	b.SetID("a30e3b91-fb77-4d22-95fa-871689c322e2")

	// shouldn't validate as empty
	_, err := b.Build()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation for 'DescriptorMap' failed on the 'required' tag")

	// Add descriptors
	err = b.AddDescriptor(Descriptor{
		ID:     "banking_input_2",
		Format: definition.CredentialFormat(definition.JWTVC),
		Path:   "$.verifiableCredential.[0]",
	})
	assert.NoError(t, err)

	err = b.AddDescriptor(Descriptor{
		ID:     "employment_input",
		Format: definition.CredentialFormat(definition.LDPVC),
		Path:   "$.verifiableCredential.[1]",
	})
	assert.NoError(t, err)

	err = b.AddDescriptor(Descriptor{
		ID:     "citizenship_input_1",
		Format: definition.CredentialFormat(definition.LDPVC),
		Path:   "$.verifiableCredential.[2]",
	})
	assert.NoError(t, err)

	presSub, err := b.Build()
	assert.NoError(t, err)
	assert.NoError(t, util.Validate(presSub))

	presSubJSON, err := util.ToJSON(presSub)
	assert.NoError(t, err)

	// get sample json from packr
	testPresSubJSON, err := testcases.GetJSONFile(testcases.SampleSubmission)
	assert.NoError(t, err)

	// Make sure our builder has the same result
	same, err := util.CompareJSON(presSubJSON, testPresSubJSON)
	assert.NoError(t, err)
	assert.True(t, same)
}
