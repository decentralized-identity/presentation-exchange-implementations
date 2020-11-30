package verifiablepresentation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission/verifiablepresentation/testcases"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

// https://identity.foundation/presentation-exchange/#presentation-submission---verifiable-presentation
func TestVerifiablePresentationBuilder(t *testing.T) {
	// First build the inner presentation submission
	b := submission.NewPresentationSubmissionBuilder("32f54163-7166-48f1-93d8-ff217bdb0653")
	b.SetID("a30e3b91-fb77-4d22-95fa-871689c322e2")

	// Add descriptors
	err := b.AddDescriptor(submission.Descriptor{
		ID:     "banking_input_2",
		Format: definition.CredentialFormat(definition.JWTVC),
		Path:   "$.verifiableCredential.[0]",
	})
	assert.NoError(t, err)

	err = b.AddDescriptor(submission.Descriptor{
		ID:     "employment_input",
		Format: definition.CredentialFormat(definition.LDPVC),
		Path:   "$.verifiableCredential.[1]",
	})
	assert.NoError(t, err)

	err = b.AddDescriptor(submission.Descriptor{
		ID:     "citizenship_input_1",
		Format: definition.CredentialFormat(definition.LDPVC),
		Path:   "$.verifiableCredential.[2]",
	})
	assert.NoError(t, err)

	presSub, err := b.Build()
	assert.NoError(t, err)
	assert.NoError(t, util.Validate(presSub))

	// Now build the Verifiable Presentation
	vpBuilder := NewVerifiablePresentationBuilder()

	// Can't build an empty presentation
	_, err = vpBuilder.Build()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation for 'Context' failed on the 'required' tag",
		"Field validation for 'Type' failed on the 'required' tag",
		"Field validation for 'PresentationSubmission' failed on the 'required' tag",
		"Field validation for 'VerifiableCredential' failed on the 'required' tag",
		"Field validation for 'Proof' failed on the 'required' tag")

	// Add the type and context
	vpBuilder.SetContext([]string{
		"https://www.w3.org/2018/credentials/v1",
		"https://identity.foundation/presentation-exchange/submission/v1",
	})
	vpBuilder.SetType([]string{"VerifiablePresentation", "PresentationSubmission"})

	// Add the presentation submission
	err = vpBuilder.SetPresentationSubmission(presSub.PresentationSubmission)
	assert.NoError(t, err)

	// Add the submission credentials
	// First cred

	vc1 := map[string]interface{}{
		"@context": "https://www.w3.org/2018/credentials/v1",
		"id":       "https://eu.com/claims/DriversLicense",
		"type": []string{
			"EUDriversLicense",
		},
		"issuer":       "did:example:123",
		"issuanceDate": "2010-01-01T19:73:24Z",
		"credentialSubject": map[string]interface{}{
			"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
			"accounts": []map[string]string{
				{
					"id":    "1234567890",
					"route": "DE-9876543210",
				},
				{
					"id":    "2457913570",
					"route": "DE-0753197542",
				},
			},
		},
	}

	// Second cred
	vc2 := map[string]interface{}{
		"@context": "https://www.w3.org/2018/credentials/v1",
		"id":       "https://eu.com/claims/DriversLicense",
		"type": []string{
			"EUDriversLicense",
		},
		"issuer":       "did:foo:123",
		"issuanceDate": "2010-01-01T19:73:24Z",
		"credentialSubject": map[string]interface{}{
			"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
			"license": map[string]string{
				"number": "34DGE352",
				"dob":    "07/13/80",
			},
		},
		"proof": map[string]string{
			"type":               "EcdsaSecp256k1VerificationKey2019",
			"created":            "2017-06-18T21:19:10Z",
			"proofPurpose":       "assertionMethod",
			"verificationMethod": "https://example.edu/issuers/keys/1",
			"jws":                "...",
		},
	}

	// Third cred
	vc3 := map[string]interface{}{
		"@context": "https://www.w3.org/2018/credentials/v1",
		"id":       "https://eu.com/claims/DriversLicense",
		"type": []string{
			"EUDriversLicense",
		},
		"issuer":       "did:foo:123",
		"issuanceDate": "2010-01-01T19:73:24Z",
		"credentialSubject": map[string]interface{}{
			"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
			"license": map[string]string{
				"number": "34DGE352",
				"dob":    "07/13/80",
			},
		},
		"proof": map[string]string{
			"type":               "RsaSignature2018",
			"created":            "2017-06-18T21:19:10Z",
			"proofPurpose":       "assertionMethod",
			"verificationMethod": "https://example.edu/issuers/keys/1",
			"jws":                "...",
		},
	}

	vpBuilder.AddVerifiableCredentials(vc1, vc2, vc3)

	// Add proof
	proof := map[string]interface{}{
		"type":               "RsaSignature2018",
		"created":            "2018-09-14T21:19:10Z",
		"proofPurpose":       "authentication",
		"verificationMethod": "did:example:ebfeb1f712ebc6f1c276e12ec21#keys-1",
		"challenge":          "1f44d55f-f161-4938-a659-f8026467f126",
		"domain":             "4jt78h47fh47",
		"jws":                "...",
	}
	vpBuilder.SetProof(proof)

	pres, err := vpBuilder.Build()
	assert.NoError(t, err)
	assert.NotEmpty(t, pres)

	presJSON, err := util.ToJSON(pres)
	assert.NoError(t, err)

	// get sample json from packr
	testVPJSON, err := testcases.GetJSONFile(testcases.SampleVerifiablePresentation)
	assert.NoError(t, err)

	// Make sure our builder has the same result
	same, err := util.CompareJSON(presJSON, testVPJSON)
	assert.NoError(t, err)
	assert.True(t, same)
}
