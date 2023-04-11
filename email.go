package secevsubid

import "encoding/json"

// EmailSubjectIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Email Identifier Format" defined in the specification.
type EmailSubjectIdentifier interface {
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

type emailSubjectIdentifier struct {
	format string
	email  string
}

func (id *emailSubjectIdentifier) Format() string {
	return id.format
}

func (id *emailSubjectIdentifier) Email() string {
	return id.email
}

func (id *emailSubjectIdentifier) Validate() error {
	if id.email == "" {
		return ErrEmptyEmail
	}

	return nil
}

func (id *emailSubjectIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat: id.Format(),
		fieldEmail:  id.Email(),
	}

	return json.Marshal(m)
}

// NewEmailSubjectIdentifier creates new instance of EmailSubjectIdentifier.,
// The argument "email" is required. If it's empty, this function raises error.
func NewEmailSubjectIdentifier(email string) (EmailSubjectIdentifier, error) {
	id := &emailSubjectIdentifier{
		format: FormatEmail,
		email:  email,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
