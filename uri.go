package secevsubid

// UriIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Uniform Resource Identifier (URI) Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-uniform-resource-identifier
type UriIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "uri".
	Format() string
	// Uri returns uri value held by the instance.
	Uri() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type uriIdentifier struct {
	F string `json:"format"`
	U string `json:"uri"`
}

func (id *uriIdentifier) Format() string {
	return id.F
}

func (id *uriIdentifier) Uri() string {
	return id.U
}

func (id *uriIdentifier) Validate() error {
	if id.U == "" {
		return ErrEmptyUri
	}

	return nil
}

// NewUriIdentifier creates new instance of UriIdentifier.,
// The argument "uri" is required. If it's empty, this function returns error.
func NewUriIdentifier(uri string) (UriIdentifier, error) {
	id := &uriIdentifier{
		F: FormatUri,
		U: uri,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
