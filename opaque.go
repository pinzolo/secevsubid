package secevsubid

import "encoding/json"

// OpaqueIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Opaque Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-opaque-identifier-format
type OpaqueIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "opaque".
	Format() string
	// Id returns id value held by the instance.
	Id() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type opaqueIdentifier struct {
	format string
	id     string
}

func (id *opaqueIdentifier) Format() string {
	return id.format
}

func (id *opaqueIdentifier) Id() string {
	return id.id
}

func (id *opaqueIdentifier) Validate() error {
	if id.id == "" {
		return ErrEmptyId
	}

	return nil
}

func (id *opaqueIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldId:     id.Id(),
	}

	return json.Marshal(m)
}

// NewOpaqueIdentifier creates new instance of OpaqueIdentifier.
// The argument "id" is required. If it's empty, this function returns error.
func NewOpaqueIdentifier(id string) (OpaqueIdentifier, error) {
	o := &opaqueIdentifier{
		format: FormatOpaque,
		id:     id,
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}

	return o, nil
}
