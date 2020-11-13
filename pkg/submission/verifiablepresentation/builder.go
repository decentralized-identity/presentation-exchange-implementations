package verifiablepresentation

import (
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/submission"
	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

type VerifiablePresentationBuilder struct {
	Presentation VerifiablePresentation
}

func NewVerifiablePresentationBuilder() *VerifiablePresentationBuilder {
	return &VerifiablePresentationBuilder{
		Presentation: VerifiablePresentation{},
	}
}

func (v *VerifiablePresentationBuilder) Build() (*VerifiablePresentation, error) {
	return &v.Presentation, util.Validate(v.Presentation)
}

func (v *VerifiablePresentationBuilder) SetContext(context []string) {
	v.Presentation.Context = context
}

func (v *VerifiablePresentationBuilder) SetType(t []string) {
	v.Presentation.Type = t
}

func (v *VerifiablePresentationBuilder) SetPresentationSubmission(p submission.PresentationSubmission) error {
	if err := util.Validate(p); err != nil {
		return err
	}
	v.Presentation.PresentationSubmission = &p
	return nil
}

func (v *VerifiablePresentationBuilder) SetProof(p interface{}) {
	v.Presentation.Proof = &p
}

func (v *VerifiablePresentationBuilder) AddVerifiableCredentials(vcs ...interface{}) {
	v.Presentation.VerifiableCredential = append(v.Presentation.VerifiableCredential, vcs...)
}
