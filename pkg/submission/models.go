package submission

import (
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/definition"
)

type PresentationSubmissionHolder struct {
	PresentationSubmission `json:"presentation_submission" validate:"required"`
}

type PresentationSubmission struct {
	ID            string       `json:"id" validate:"required"`
	DefinitionID  string       `json:"definition_id" validate:"required"`
	Locale        string       `json:"locale,omitempty"`
	DescriptorMap []Descriptor `json:"descriptor_map" validate:"required"`
}

type Descriptor struct {
	ID     string                      `json:"id" validate:"required"`
	Path   string                      `json:"path" validate:"required"`
	Format definition.CredentialFormat `json:"format" validate:"required"`
}
