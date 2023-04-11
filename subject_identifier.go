package secevsubid

import (
	"encoding/json"
	"fmt"
)

// SubjectIdentifier is interface for handling transparently each identifier formats defined at the specification of Subject Identifiers for Security Event Tokens.
// See: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers.
type SubjectIdentifier interface {
	// Format returns name of the format actually held by the instance.
	Format() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

// Wrapper internally holds a single instance of SubjectIdentifier.
// Since json.Marshal cannot assign values to interface, it can be deserialized from JSON to a SubjectIdentifier instance dynamically via Wrapper.
type Wrapper struct {
	v SubjectIdentifier
}

// Value returns the instance of SubjectIdentifier held internally.
func (w *Wrapper) Value() SubjectIdentifier {
	return w.v
}

// MarshalJSON implements json.Marshaler.
// Returns JSON representation of the SubjectIdentifier held internally.
func (w *Wrapper) MarshalJSON() ([]byte, error) {
	return w.v.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler
func (w *Wrapper) UnmarshalJSON(b []byte) error {
	id, err := DecodeJSON(b)
	if err != nil {
		return err
	}

	w.v = id
	return nil
}

// NewWrapper creates new instance of Wrapper.
func NewWrapper(id SubjectIdentifier) *Wrapper {
	return &Wrapper{v: id}
}

// DecodeJSON decodes to the appropriate SubjectIdentifier instance.
// Which format is decoded is determined by the value of the "format" field,
// and an error is returned if there is no corresponding format.
func DecodeJSON(b []byte) (SubjectIdentifier, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	f, ok := m[fieldFormat].(string)
	if !ok {
		return nil, ErrNoFormat
	}

	fn := func(name string) string {
		return m[name].(string)
	}
	switch f {
	case FormatAccount:
		return NewAccountIdentifier(fn(fieldUri))
	case FormatEmail:
		return NewEmailIdentifier(fn(fieldEmail))
	case FormatIssuerSubject:
		return NewIssuerSubjectIdentifier(fn(fieldIssuer), fn(fieldSubject))
	case FormatOpaque:
		return NewOpaqueIdentifier(fn(fieldId))
	case FormatPhoneNumber:
		return NewPhoneNumberIdentifier(fn(fieldPhoneNumber))
	case FormatDid:
		return NewDidIdentifier(fn(fieldUrl))
	case FormatUri:
		return NewUriIdentifier(fn(FormatUri))
	case FormatAliases:
		// TODO
		return nil, nil
	}

	return nil, fmt.Errorf("unknown format: %s", f)
}
