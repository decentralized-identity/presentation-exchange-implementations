package submission

import (
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

type PresentationSubmissionBuilder struct {
	Submission PresentationSubmission
}

func NewPresentationSubmissionBuilder() *PresentationSubmissionBuilder {
	return &PresentationSubmissionBuilder{
		Submission: PresentationSubmission{},
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

func (p *PresentationSubmissionBuilder) SetLocale(locale string) {
	p.Submission.Locale = locale
}
