package secevsubid

import "encoding/json"

// IssuerSubjectIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Issuer and Subject Identifier Format" defined in the specification.
type IssuerSubjectIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "issSub".
	Format() string
	// Issuer returns issuer value held by the instance.
	Issuer() string
	// Subject returns subject value held by the instance.
	Subject() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}

type issSubIdentifier struct {
	format  string
	issuer  string
	subject string
}

func (id *issSubIdentifier) Format() string {
	return id.format
}

func (id *issSubIdentifier) Issuer() string {
	return id.issuer
}

func (id *issSubIdentifier) Subject() string {
	return id.subject
}

func (id *issSubIdentifier) Validate() error {
	if id.issuer == "" {
		return ErrEmptyIssuer
	}

	if id.subject == "" {
		return ErrEmptySubject
	}

	return nil
}

func (id *issSubIdentifier) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		fieldFormat:  id.Format(),
		fieldIssuer:  id.Issuer(),
		fieldSubject: id.Subject(),
	}

	return json.Marshal(m)
}

// NewIssuerSubjectIdentifier creates new instance of IssuerSubjectIdentifier.
// The argument "issuer" and "subject" is required. If either one of them is empty, this function returns error.
func NewIssuerSubjectIdentifier(issuer string, subject string) (IssuerSubjectIdentifier, error) {
	id := &issSubIdentifier{
		format:  FormatIssuerSubject,
		issuer:  issuer,
		subject: subject,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
