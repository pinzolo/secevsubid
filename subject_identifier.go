package secevsubid

// SubjectIdentifier is interface for handling transparently each identifier formats defined at Subject Identifiers for Security Event Tokens.
// See: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers.
type SubjectIdentifier interface {
	// Format returns name of the format actually held by the instance.
	Format() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
	// MarshalJSON is required for instance to be converted to JSON.
	MarshalJSON() ([]byte, error)
}
