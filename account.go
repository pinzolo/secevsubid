package secevsubid

// AccountIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Account Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-account-identifier-format
type AccountIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "account".
	Format() Format
	// Uri returns uri value held by the instance.
	Uri() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type accountIdentifier struct {
	F Format `json:"format"`
	U string `json:"uri"`
}

func (id *accountIdentifier) Format() Format {
	return id.F
}

func (id *accountIdentifier) Uri() string {
	return id.U
}

func (id *accountIdentifier) Validate() error {
	if id.U == "" {
		return ErrEmptyUri
	}

	return nil
}

// NewAccountIdentifier creates new instance of AccountIdentifier.
// The argument "uri" is required. If it's empty, this function returns error.
func NewAccountIdentifier(uri string) (AccountIdentifier, error) {
	id := &accountIdentifier{
		F: FormatAccount,
		U: uri,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
