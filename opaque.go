package secevsubid

// OpaqueIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Opaque Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-opaque-identifier-format
type OpaqueIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "opaque".
	Format() Format
	// Id returns id value held by the instance.
	Id() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type opaqueIdentifier struct {
	F Format `json:"format"`
	I string `json:"id"`
}

func (id *opaqueIdentifier) Format() Format {
	return id.F
}

func (id *opaqueIdentifier) Id() string {
	return id.I
}

func (id *opaqueIdentifier) Validate() error {
	if id.I == "" {
		return ErrEmptyId
	}

	return nil
}

// NewOpaqueIdentifier creates new instance of OpaqueIdentifier.
// The argument "id" is required. If it's empty, this function returns error.
func NewOpaqueIdentifier(id string) (OpaqueIdentifier, error) {
	o := &opaqueIdentifier{
		F: FormatOpaque,
		I: id,
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}

	return o, nil
}
