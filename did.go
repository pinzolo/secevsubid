package secevsubid

import "encoding/json"

// DidIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Decentralized Identifier (DID) Format" defined in the specification.
type DidIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "did".
	Format() string
	// Url returns url value held by the instance.
	Url() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type didIdentifier struct {
	format string
	url    string
}

func (id *didIdentifier) Format() string {
	return id.format
}

func (id *didIdentifier) Url() string {
	return id.url
}

func (id *didIdentifier) Validate() error {
	if id.url == "" {
		return ErrEmptyUrl
	}

	return nil
}

func (id *didIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldUrl:    id.Url(),
	}

	return json.Marshal(m)
}

// NewDidIdentifier creates new instance of DidIdentifier.,
// The argument "url" is required. If it's empty, this function returns error.
func NewDidIdentifier(url string) (DidIdentifier, error) {
	id := &didIdentifier{
		format: FormatDid,
		url:    url,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
