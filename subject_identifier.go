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
	return json.Marshal(w.v)
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

var extractStringValue = func(m map[string]interface{}, name string) string {
	s, ok := m[name]
	if !ok {
		return ""
	}
	return s.(string)
}

// DecodeJSON decodes to the appropriate SubjectIdentifier instance.
// Which format is decoded is determined by the value of the "format" field,
// and an error is returned if there is no corresponding format.
func DecodeJSON(b []byte) (SubjectIdentifier, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return decodeIdentifier(m)
}

func decodeIdentifier(m map[string]interface{}) (SubjectIdentifier, error) {
	f, ok := m[fieldFormat].(string)
	if !ok {
		return nil, ErrNoFormat
	}

	switch f {
	case FormatAccount:
		return NewAccountIdentifier(extractStringValue(m, fieldUri))
	case FormatEmail:
		return NewEmailIdentifier(extractStringValue(m, fieldEmail))
	case FormatIssuerSubject:
		return NewIssuerSubjectIdentifier(extractStringValue(m, fieldIssuer), extractStringValue(m, fieldSubject))
	case FormatOpaque:
		return NewOpaqueIdentifier(extractStringValue(m, fieldId))
	case FormatPhoneNumber:
		return NewPhoneNumberIdentifier(extractStringValue(m, fieldPhoneNumber))
	case FormatDid:
		return NewDidIdentifier(extractStringValue(m, fieldUrl))
	case FormatUri:
		return NewUriIdentifier(extractStringValue(m, fieldUri))
	case FormatAliases:
		return decodeAliases(m)
	}

	return nil, fmt.Errorf("unknown format: %s", f)
}

func decodeAliases(m map[string]interface{}) (SubjectIdentifier, error) {
	vs := m[fieldIdentifiers].([]interface{})
	if len(vs) == 0 {
		return nil, ErrEmptyIdentifiers
	}

	ids := make([]SubjectIdentifier, len(vs))
	for i, v := range vs {
		d, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("not JSON object: %v", v)
		}

		f := extractStringValue(d, fieldFormat)
		if f == "" {
			return nil, ErrNoFormat
		}
		if f == FormatAliases {
			return nil, ErrNestedAliases
		}

		id, err := decodeIdentifier(d)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	id, err := NewAliasesIdentifier(ids...)
	if err != nil {
		return nil, err
	}

	return id, nil
}
