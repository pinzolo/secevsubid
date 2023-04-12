package secevsubid

import "encoding/json"

// AccountIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Account Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-account-identifier-format
type AccountIdentifier interface {
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

type accountIdentifier struct {
	format string
	uri    string
}

func (id *accountIdentifier) Format() string {
	return id.format
}

func (id *accountIdentifier) Uri() string {
	return id.uri
}

func (id *accountIdentifier) Validate() error {
	if id.uri == "" {
		return ErrEmptyUri
	}

	return nil
}

func (id *accountIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldUri:    id.Uri(),
	}

	return json.Marshal(m)
}

// NewAccountIdentifier creates new instance of AccountIdentifier.
// The argument "uri" is required. If it's empty, this function returns error.
func NewAccountIdentifier(uri string) (AccountIdentifier, error) {
	id := &accountIdentifier{
		format: FormatAccount,
		uri:    uri,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
