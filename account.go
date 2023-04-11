package secevsubid

import "encoding/json"

// AccountSubjectIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Account Identifier Format" defined in the specification.
type AccountSubjectIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "account".
	Format() string
	// Uri returns uri value held by the instance.
	Uri() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type accountSubjectIdentifier struct {
	format string
	uri    string
}

func (id *accountSubjectIdentifier) Format() string {
	return id.format
}

func (id *accountSubjectIdentifier) Uri() string {
	return id.uri
}

func (id *accountSubjectIdentifier) Validate() error {
	if id.uri == "" {
		return ErrEmptyUri
	}

	return nil
}

func (id *accountSubjectIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldUri:    id.Uri(),
	}

	return json.Marshal(m)
}

// NewAccountSubjectIdentifier creates new instance of AccountSubjectIdentifier.,
// The argument "uri" is required. If it's empty, this function raises error.
func NewAccountSubjectIdentifier(uri string) (AccountSubjectIdentifier, error) {
	id := &accountSubjectIdentifier{
		format: FormatAccount,
		uri:    uri,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
