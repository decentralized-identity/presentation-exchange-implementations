package verifiablepresentation

import (
	"go.wday.io/credentials-open-source/presentation-exchange/pkg/submission"
)

type VerifiablePresentation struct {
	Context                []string                           `json:"@context" validate:"required"`
	Type                   []string                           `json:"type" validate:"required"`
	PresentationSubmission *submission.PresentationSubmission `json:"presentation_submission" validate:"required"`
	VerifiableCredential   []interface{}                      `json:"verifiableCredential" validate:"required"`
	Proof                  interface{}                        `json:"proof"`
}
