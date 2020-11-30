package submission

import (
	"github.com/google/uuid"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

type PresentationSubmissionBuilder struct {
	Submission PresentationSubmission
}

func NewPresentationSubmissionBuilder(definitionID string) *PresentationSubmissionBuilder {
	return &PresentationSubmissionBuilder{
		Submission: PresentationSubmission{
			ID:           uuid.New().String(),
			DefinitionID: definitionID,
		},
	}
}

func (p *PresentationSubmissionBuilder) Build() (*PresentationSubmissionHolder, error) {
	return &PresentationSubmissionHolder{
		PresentationSubmission: p.Submission,
	}, util.Validate(p.Submission)
}

func (p *PresentationSubmissionBuilder) AddDescriptor(d Descriptor) error {
	if err := util.Validate(d); err != nil {
		return err
	}
	p.Submission.DescriptorMap = append(p.Submission.DescriptorMap, d)
	return nil
}

func (p *PresentationSubmissionBuilder) SetID(id string) {
	p.Submission.ID = id
}

func (p *PresentationSubmissionBuilder) SetDefinitionID(definitionID string) {
	p.Submission.DefinitionID = definitionID
}

func (p *PresentationSubmissionBuilder) SetLocale(locale string) {
	p.Submission.Locale = locale
}
