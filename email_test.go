package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestEmailSubjectIdentifier(t *testing.T) {
	wantEmail := "user@example.com"
	id, err := secevsubid.NewEmailSubjectIdentifier(wantEmail)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatEmail {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatEmail)
	}
	if id.Email() != wantEmail {
		t.Errorf("invalid email: got = %s, want = %s", id.Email(), wantEmail)
	}

	wantJSON := fmt.Sprintf(`{"email":"%s","format":"email"}`, wantEmail)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestEmailSubjectIdentifierWithEmptyUri(t *testing.T) {
	_, err := secevsubid.NewEmailSubjectIdentifier("")
	if err == nil {
		t.Error("error should be raised when email is empty")
	}
}
