package verifiablepresentation

import (
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission"
)

type VerifiablePresentation struct {
	Context                []string                           `json:"@context" validate:"required"`
	Type                   []string                           `json:"type" validate:"required"`
	PresentationSubmission *submission.PresentationSubmission `json:"presentation_submission" validate:"required"`
	VerifiableCredential   []interface{}                      `json:"verifiableCredential" validate:"required"`
	Proof                  interface{}                        `json:"proof"`
}
