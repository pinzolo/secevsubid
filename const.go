package secevsubid

import "errors"

const (
	// FormatAccount is the format name for Account Identifier Format.
	FormatAccount = "account"
	// FormatEmail is the format name for Email Identifier Format.
	FormatEmail = "email"
	// FormatIssuerSubject is the format name for Issuer and Subject Identifier Format.
	FormatIssuerSubject = "iss_sub"
	// FormatOpaque is the format name for Opaque Identifier Format.
	FormatOpaque = "opaque"
	// FormatPhoneNumber is the format name for Phone Number Identifier Format.
	FormatPhoneNumber = "phone_number"
	// FormatDid is the format name for Decentralized Identifier (DID) Format.
	FormatDid = "did"
	// FormatUri is the format name for Uniform Resource Identifier (URI) Format.
	FormatUri = "uri"
	// FormatAliases is the format name for Aliases Identifier Format.
	FormatAliases = "aliases"

	// FieldFormat is the field name for "format" field.
	// This field is used in all identifier format.
	fieldFormat = "format"
	// FieldUri is the field name for "uri" field.
	// This field is used in Account Identifier Format, Uniform Resource Identifier (URI) Format.
	fieldUri = "uri"
	// FieldEmail is the field name for "email" field.
	// This field is used in Email Identifier Format.
	fieldEmail = "email"
	// FieldIssuer is the field name for "iss" field.
	// This field is used in Issuer and Subject Identifier Format.
	fieldIssuer = "iss"
	// FieldSubject is the field name for "sub" field.
	// This field is used in Issuer and Subject Identifier Format.
	fieldSubject = "sub"
	// FieldId is the field name for "id" field.
	// This field is used in Opaque Identifier Format.
	fieldId = "id"
	// FieldPhoneNumber is the field name for "phoneNumber" field.
	// This field is used in Phone Number Identifier Format.
	fieldPhoneNumber = "phone_number"
	// FieldUrl is the field name for "url" field.
	// This field is used in Decentralized Identifier (DID) Identifier Format.
	fieldUrl = "url"
	// FieldIdentifiers is the field name for "identifiers" field.
	// This field is used in Aliases Identifier Format
	fieldIdentifiers = "identifiers"
)

var (
	// ErrEmptyUri is error raised when email value does not exist at generation time.
	ErrEmptyUri = errors.New("empty uri")
	// ErrEmptyEmail is error raised when email value does not exist at generation time.
	ErrEmptyEmail = errors.New("empty email")
	// ErrEmptyIssuer is error raised when iss value does not exist at generation time.
	ErrEmptyIssuer = errors.New("empty iss")
	// ErrEmptySubject is error raised when sub value does not exist at generation time.
	ErrEmptySubject = errors.New("empty sub")
	// ErrEmptyId is error raised when id value does not exist at generation time.
	ErrEmptyId = errors.New("empty id")
	// ErrEmptyPhoneNumber is error raised when phoneNumber value does not exist at generation time.
	ErrEmptyPhoneNumber = errors.New("empty phone number")
	// ErrEmptyUrl is error raised when url value does not exist at generation time.
	ErrEmptyUrl = errors.New("empty url")
)
