package secevsubid

import "errors"

const (
	// FormatAccount is the format name for Account Identifier Format.
	FormatAccount = "account"
	// FormatEmail is the format name for Email Identifier Format.
	FormatEmail = "email"

	// FieldFormat is the field name for "format" field.
	// This field is used in all identifier format.
	fieldFormat = "format"
	// FieldUri is the field name for "uri" field.
	// This field is used in Account Identifier Format, Uniform Resource Identifier (URI) Format.
	fieldUri = "uri"
)

var (
	// ErrEmptyUri is error raised when email value does not exist at generation time.
	ErrEmptyUri = errors.New("empty email")
	// ErrEmptyEmail is error raised when email value does not exist at generation time.
	ErrEmptyEmail = errors.New("empty email")
)
