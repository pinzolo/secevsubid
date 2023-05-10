package secevsubid

// DidIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Decentralized Identifier (DID) Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-decentralized-identifier-di
type DidIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "did".
	Format() Format
	// Url returns url value held by the instance.
	Url() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type didIdentifier struct {
	F Format `json:"format"`
	U string `json:"url"`
}

func (id *didIdentifier) Format() Format {
	return id.F
}

func (id *didIdentifier) Url() string {
	return id.U
}

func (id *didIdentifier) Validate() error {
	if id.U == "" {
		return ErrEmptyUrl
	}

	return nil
}

// NewDidIdentifier creates new instance of DidIdentifier.
// The argument "url" is required. If it's empty, this function returns error.
func NewDidIdentifier(url string) (DidIdentifier, error) {
	id := &didIdentifier{
		F: FormatDid,
		U: url,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
