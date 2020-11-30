package definition

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

type PresentationDefinitionBuilder struct {
	Definition PresentationDefinition
}

func NewPresentationDefinitionBuilder() *PresentationDefinitionBuilder {
	return &PresentationDefinitionBuilder{
		Definition: PresentationDefinition{
			ID: uuid.New().String(),
		},
	}
}

func (p *PresentationDefinitionBuilder) Build() (*PresentationDefinitionHolder, error) {
	return &PresentationDefinitionHolder{
		PresentationDefinition: p.Definition,
	}, util.Validate(p.Definition)
}

func (p *PresentationDefinitionBuilder) SetName(name string) {
	p.Definition.Name = name
}

func (p *PresentationDefinitionBuilder) SetID(id string) {
	p.Definition.ID = id
}

func (p *PresentationDefinitionBuilder) SetPurpose(purpose string) {
	p.Definition.Purpose = purpose
}

func (p *PresentationDefinitionBuilder) SetLocale(locale string) {
	p.Definition.Locale = locale
}

// Submission Requirement Builders //

func (p *PresentationDefinitionBuilder) SetJWTFormat(format JWTFormat, algs []string) error {
	if len(algs) < 1 {
		return fmt.Errorf("must set one or more algs for the jwt type<%s>", format)
	}
	if p.Definition.Format == nil {
		p.Definition.Format = &Format{}
	}
	switch format {
	case JWT:
		p.Definition.Format.JWT = &JWTType{Alg: algs}
	case JWTVC:
		p.Definition.Format.JWTVC = &JWTType{Alg: algs}
	case JWTVP:
		p.Definition.Format.JWTVP = &JWTType{Alg: algs}
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
	return nil
}

func (p *PresentationDefinitionBuilder) SetLDPFormat(format LDPFormat, proofTypes []string) error {
	if len(proofTypes) < 1 {
		return fmt.Errorf("must set one or more proof types for the ldp type<%s>", format)
	}
	if p.Definition.Format == nil {
		p.Definition.Format = &Format{}
	}
	switch format {
	case LDP:
		p.Definition.Format.LDP = &LDPType{ProofType: proofTypes}
	case LDPVC:
		p.Definition.Format.LDPVC = &LDPType{ProofType: proofTypes}
	case LDPVP:
		p.Definition.Format.LDPVP = &LDPType{ProofType: proofTypes}
	default:
		return fmt.Errorf("unknown format: %s", format)
	}
	return nil
}

func (p *PresentationDefinitionBuilder) AddSubmissionRequirements(srs ...SubmissionRequirement) error {
	var errs *multierror.Error
	for _, sr := range srs {
		if err := validateSubmissionRequirement(sr); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	if errs == nil {
		p.Definition.SubmissionRequirements = srs
	}
	if errs == nil {
		return nil
	}
	return errs.ErrorOrNil()
}

func validateSubmissionRequirement(sr SubmissionRequirement) error {
	if sr.From == "" && len(sr.FromNested) == 0 {
		return errors.New("from must have a value")
	}
	if sr.From != "" && len(sr.FromNested) > 0 {
		return errors.New("from and from_nested are mutually exclusive fields")
	}
	return util.Validate(sr)
}

// Input Descriptor Builders //

func (p *PresentationDefinitionBuilder) AddInputDescriptor(i InputDescriptor) error {
	if err := util.Validate(i); err != nil {
		return err
	}
	p.Definition.InputDescriptors = append(p.Definition.InputDescriptors, i)
	return nil
}

func NewInputDescriptor(id, name, purpose, metadata string) *InputDescriptor {
	return &InputDescriptor{
		ID:       id,
		Name:     name,
		Purpose:  purpose,
		Metadata: metadata,
	}
}

func (i *InputDescriptor) AddSchema(s Schema) error {
	if err := util.Validate(s); err != nil {
		return err
	}
	i.Schema = append(i.Schema, s)
	return nil
}

func (i *InputDescriptor) SetConstraints(fields ...Field) error {
	for _, f := range fields {
		if f.Predicate != nil && f.Filter == nil {
			return fmt.Errorf("field cannot have a predicate preference without a filter: %+v", f)
		}
	}
	if i.Constraints == nil {
		i.Constraints = &Constraints{
			Fields: fields,
		}
	} else {
		i.Constraints.Fields = fields
	}
	return util.Validate(i.Constraints)
}

func (i *InputDescriptor) SetSubjectIsIssuer(preference Preference) error {
	if i.Constraints == nil {
		i.Constraints = &Constraints{}
	}
	i.Constraints.SubjectIsIssuer = &preference
	return nil
}

func (i *InputDescriptor) SetSubjectIsHolder(preference Preference) error {
	if i.Constraints == nil {
		i.Constraints = &Constraints{}
	}
	i.Constraints.SubjectIsHolder = &preference
	return nil
}

func (i *InputDescriptor) SetConstraintsLimitDisclosure(limitDisclosure bool) {
	if i.Constraints == nil {
		i.Constraints = &Constraints{
			LimitDisclosure: limitDisclosure,
		}
	}
	i.Constraints.LimitDisclosure = limitDisclosure
}

func NewConstraintsField(path []string) *Field {
	return &Field{Path: path}
}

func (f *Field) SetPurpose(purpose string) {
	f.Purpose = purpose
}

func (f *Field) SetFilter(filter Filter) error {
	if err := util.Validate(filter); err != nil {
		return err
	}
	f.Filter = &filter
	return nil
}
