package definition

type (
	CredentialFormat string
	JWTFormat        CredentialFormat
	LDPFormat        CredentialFormat
)

const (
	JWT   JWTFormat = "jwt"
	JWTVC JWTFormat = "jwt_vc"
	JWTVP JWTFormat = "jwt_vp"

	LDP   LDPFormat = "ldp"
	LDPVC LDPFormat = "ldp_vc"
	LDPVP LDPFormat = "ldp_vp"
)

type PresentationDefinitionHolder struct {
	PresentationDefinition `json:"presentation_definition" validate:"required"`
}

type PresentationDefinition struct {
	Name                   string                  `json:"name,omitempty"`
	Purpose                string                  `json:"purpose,omitempty"`
	Locale                 string                  `json:"locale,omitempty"`
	Format                 *Format                 `json:"format,omitempty"`
	SubmissionRequirements []SubmissionRequirement `json:"submission_requirements,omitempty"`
	InputDescriptors       []InputDescriptor       `json:"input_descriptors" validate:"required"`
}

type Format struct {
	JWT   *JWTType `json:"jwt,omitempty"`
	JWTVC *JWTType `json:"jwt_vc,omitempty"`
	JWTVP *JWTType `json:"jwt_vp,omitempty"`

	LDP   *LDPType `json:"ldp,omitempty"`
	LDPVC *LDPType `json:"ldp_vc,omitempty"`
	LDPVP *LDPType `json:"ldp_vp,omitempty"`
}

type JWTType struct {
	Alg []string `json:"alg,omitempty" validate:"required"`
}

type LDPType struct {
	ProofType []string `json:"proof_type,omitempty" validate:"required"`
}

type SubmissionRequirement struct {
	Name    string `json:"name,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Rule    string `json:"rule" validate:"required"`
	Count   int    `json:"count,omitempty" validate:"min=1"`
	Minimum int    `json:"min,omitempty"`
	Maximum int    `json:"max,omitempty"`

	// Either an array of SubmissionRequirement or a string value
	FromOption `validate:"required"`
}

type FromOption struct {
	From       string                  `json:"from,omitempty"`
	FromNested []SubmissionRequirement `json:"from_nested,omitempty"`
}

type InputDescriptor struct {
	ID          string       `json:"id,omitempty" validate:"required"`
	Group       []string     `json:"group,omitempty"`
	Schema      *Schema      `json:"schema,omitempty" validate:"required"`
	Constraints *Constraints `json:"constraints,omitempty"`
}

type Schema struct {
	URI      []string `json:"uri,omitempty" validate:"required,min=1"`
	Name     string   `json:"name,omitempty" validate:"required"`
	Purpose  string   `json:"purpose,omitempty"`
	Metadata string   `json:"metadata,omitempty"`
}

type Constraints struct {
	LimitDisclosure bool    `json:"limit_disclosure,omitempty"`
	Fields          []Field `json:"fields,omitempty"`
}

type Field struct {
	Path    []string `json:"path,omitempty" validate:"required"`
	Purpose string   `json:"purpose,omitempty"`
	Filter  *Filter  `json:"filter,omitempty"`
}

type Filter struct {
	Type      string `json:"type" validate:"required"`
	Format    string `json:"format,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Minimum   string `json:"minimum,omitempty"`
	MinLength int    `json:"minLength,omitempty"`
	MaxLength int    `json:"maxLength,omitempty"`
}
