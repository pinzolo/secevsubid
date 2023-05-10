package secevsubid

// EmailIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Email Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-email-identifier-format
type EmailIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "email".
	Format() Format
	// Email returns email value held by the instance.
	Email() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type emailIdentifier struct {
	F Format `json:"format"`
	E string `json:"email"`
}

func (id *emailIdentifier) Format() Format {
	return id.F
}

func (id *emailIdentifier) Email() string {
	return id.E
}

func (id *emailIdentifier) Validate() error {
	if id.E == "" {
		return ErrEmptyEmail
	}

	return nil
}

// NewEmailIdentifier creates new instance of EmailIdentifier.
// The argument "email" is required. If it's empty, this function returns error.
func NewEmailIdentifier(email string) (EmailIdentifier, error) {
	id := &emailIdentifier{
		F: FormatEmail,
		E: email,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
