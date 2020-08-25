package submission

import (
	"go.wday.io/credentials-open-source/presentation-exchange/pkg/definition"
)

type PresentationSubmissionHolder struct {
	PresentationSubmission `json:"presentation_submission" validate:"required"`
}

type PresentationSubmission struct {
	Locale        string       `json:"locale,omitempty"`
	DescriptorMap []Descriptor `json:"descriptor_map" validate:"required"`
}

type Descriptor struct {
	ID     string                      `json:"id" validate:"required"`
	Path   string                      `json:"path" validate:"required"`
	Format definition.CredentialFormat `json:"format" validate:"required"`
}
