package secevsubid

import (
	"reflect"
)

// AliasesIdentifier is one of the sub-interfaces of SubjectIdentifier.
// It represents the "Aliases Identifier Format" defined in the specification.
// Reference: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers#name-aliases-identifier-format
type AliasesIdentifier interface {
	// Format returns name of the format actually held by the instance.
	// The value is the fixed value "aliases".
	Format() string
	// Identifiers returns SubjectIdentifier list value held by the instance.
	Identifiers() []SubjectIdentifier
	// ContainsIdentifier returns whether a SubjectIdentifier with the same content as the argument already exists.
	ContainsIdentifier(identifier SubjectIdentifier) bool
	// AddIdentifier adds new SubjectIdentifier to internal list.
	// In the following cases this method returns an error.
	//   * The argument is in Aliases Identifier Format.
	//   * A SubjectIdentifier with the same content as the argument already exists.
	AddIdentifier(identifier SubjectIdentifier) error
	// Validate values held and returns an error if there is a problem.
	Validate() error
}

type aliasesIdentifier struct {
	F   string              `json:"format"`
	Ids []SubjectIdentifier `json:"identifiers"`
}

func (id *aliasesIdentifier) Format() string {
	return id.F
}

func (id *aliasesIdentifier) Identifiers() []SubjectIdentifier {
	c := make([]SubjectIdentifier, len(id.Ids))
	_ = copy(c, id.Ids)
	return c
}

func (id *aliasesIdentifier) Validate() error {
	if len(id.Ids) == 0 {
		return ErrEmptyIdentifiers
	}

	return nil
}

func (id *aliasesIdentifier) ContainsIdentifier(identifier SubjectIdentifier) bool {
	for _, v := range id.Ids {
		if v.Format() == identifier.Format() && reflect.DeepEqual(v, identifier) {
			return true
		}
	}

	return false
}

func (id *aliasesIdentifier) AddIdentifier(identifier SubjectIdentifier) error {
	if identifier.Format() == FormatAliases {
		return ErrNestedAliases
	}

	if id.ContainsIdentifier(identifier) {
		return ErrDuplicatedIdentifier
	}

	id.Ids = append(id.Ids, identifier)
	return nil
}

// NewAliasesIdentifier creates new instance of AliasesIdentifier.
func NewAliasesIdentifier(identifiers ...SubjectIdentifier) (AliasesIdentifier, error) {
	id := &aliasesIdentifier{F: FormatAliases}

	if len(identifiers) > 0 {
		for _, i := range identifiers {
			err := id.AddIdentifier(i)
			if err != nil {
				return nil, err
			}
		}
	}

	return id, nil
}
