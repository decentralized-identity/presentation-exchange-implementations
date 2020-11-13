package submission

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decentralized-identity/presentation-exchange-implementations/internal/schema"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission/verifiablepresentation/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

func TestPresentationSubmission(t *testing.T) {
	samplePresSub, err := testcases.GetJSONFile(testcases.SampleSubmission)
	assert.NoError(t, err)

	var presSub PresentationSubmissionHolder
	presSubBytes := []byte(samplePresSub)
	err = json.Unmarshal(presSubBytes, &presSub)
	assert.NoError(t, err)

	assert.NoError(t, util.Validate(presSub))

	// Roundtrip and compare
	roundTripBytes, err := json.Marshal(presSub)
	assert.NoError(t, err)

	true, err := schema.JSONBytesEqual(presSubBytes, roundTripBytes)
	assert.NoError(t, err)
	assert.True(t, true)
}
