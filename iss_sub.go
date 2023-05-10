package secevsubid

// IssuerSubjectIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Issuer and Subject Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-issuer-and-subject-identifi
type IssuerSubjectIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "issSub".
	Format() Format
	// Issuer returns issuer value held by the instance.
	Issuer() string
	// Subject returns subject value held by the instance.
	Subject() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type issSubIdentifier struct {
	F Format `json:"format"`
	I string `json:"iss"`
	S string `json:"sub"`
}

func (id *issSubIdentifier) Format() Format {
	return id.F
}

func (id *issSubIdentifier) Issuer() string {
	return id.I
}

func (id *issSubIdentifier) Subject() string {
	return id.S
}

func (id *issSubIdentifier) Validate() error {
	if id.I == "" {
		return ErrEmptyIssuer
	}

	if id.S == "" {
		return ErrEmptySubject
	}

	return nil
}

// NewIssuerSubjectIdentifier creates new instance of IssuerSubjectIdentifier.
// The argument "issuer" and "subject" is required. If either one of them is empty, this function returns error.
func NewIssuerSubjectIdentifier(issuer string, subject string) (IssuerSubjectIdentifier, error) {
	id := &issSubIdentifier{
		F: FormatIssuerSubject,
		I: issuer,
		S: subject,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
