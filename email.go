package secevsubid

import "encoding/json"

// EmailIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Email Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-email-identifier-format
type EmailIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "email".
	Format() string
	// Email returns email value held by the instance.
	Email() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type emailIdentifier struct {
	format string
	email  string
}

func (id *emailIdentifier) Format() string {
	return id.format
}

func (id *emailIdentifier) Email() string {
	return id.email
}

func (id *emailIdentifier) Validate() error {
	if id.email == "" {
		return ErrEmptyEmail
	}

	return nil
}

func (id *emailIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldEmail:  id.Email(),
	}

	return json.Marshal(m)
}

// NewEmailIdentifier creates new instance of EmailIdentifier.
// The argument "email" is required. If it's empty, this function returns error.
func NewEmailIdentifier(email string) (EmailIdentifier, error) {
	id := &emailIdentifier{
		format: FormatEmail,
		email:  email,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
