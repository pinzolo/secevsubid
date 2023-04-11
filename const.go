package secevsubid

import "errors"

const (
	// FormatAccount is the format name for Account Identifier Format.
	FormatAccount = "account"

	// FieldFormat is the field name for "format" field.
	// This field is used in all identifier format.
	fieldFormat = "format"
	// FieldUri is the field name for "uri" field.
	// This field is used in Account Identifier Format, Uniform Resource Identifier (URI) Format.
	fieldUri = "uri"
)

var (
	// ErrEmptyUri is error raised when uri value does not exist at generation time.
	ErrEmptyUri = errors.New("empty uri")
)
