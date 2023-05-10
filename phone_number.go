package secevsubid

// PhoneNumberIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Phone Number Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-phone-number-identifier-for
type PhoneNumberIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "phoneNumber".
	Format() Format
	// PhoneNumber returns phoneNumber value held by the instance.
	PhoneNumber() string
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type phoneNumberIdentifier struct {
	F Format `json:"format"`
	N string `json:"phone_number"`
}

func (id *phoneNumberIdentifier) Format() Format {
	return id.F
}

func (id *phoneNumberIdentifier) PhoneNumber() string {
	return id.N
}

func (id *phoneNumberIdentifier) Validate() error {
	if id.N == "" {
		return ErrEmptyPhoneNumber
	}

	return nil
}

// NewPhoneNumberIdentifier creates new instance of PhoneNumberIdentifier.
// The argument "phoneNumber" is required. If it's empty, this function returns error.
func NewPhoneNumberIdentifier(phoneNumber string) (PhoneNumberIdentifier, error) {
	id := &phoneNumberIdentifier{
		F: FormatPhoneNumber,
		N: phoneNumber,
	}
	if err := id.Validate(); err != nil {
		return nil, err
	}

	return id, nil
}
