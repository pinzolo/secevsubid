package secevsubid

import "encoding/json"

// PhoneNumberIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Phone Number Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-phone-number-identifier-for
type PhoneNumberIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "phoneNumber".
	Format() string
	// PhoneNumber returns phoneNumber value held by the instance.
	PhoneNumber() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type phoneNumberIdentifier struct {
	format      string
	phoneNumber string
}

func (id *phoneNumberIdentifier) Format() string {
	return id.format
}

func (id *phoneNumberIdentifier) PhoneNumber() string {
	return id.phoneNumber
}

func (id *phoneNumberIdentifier) Validate() error {
	if id.phoneNumber == "" {
		return ErrEmptyPhoneNumber
	}

	return nil
}

func (id *phoneNumberIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat:      id.Format(),
		fieldPhoneNumber: id.PhoneNumber(),
	}

	return json.Marshal(m)
}

// NewPhoneNumberIdentifier creates new instance of PhoneNumberIdentifier.
// The argument "phoneNumber" is required. If it's empty, this function returns error.
func NewPhoneNumberIdentifier(phoneNumber string) (PhoneNumberIdentifier, error) {
	id := &phoneNumberIdentifier{
		format:      FormatPhoneNumber,
		phoneNumber: phoneNumber,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
