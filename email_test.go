package secevsubid_test

import (
	"encoding/json"
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestEmailIdentifier(t *testing.T) {
	wantEmail := "user@example.com"
	id, err := secevsubid.NewEmailIdentifier(wantEmail)
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
	b, err := json.Marshal(id)
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestEmailIdentifierWithEmptyEmail(t *testing.T) {
	_, err := secevsubid.NewEmailIdentifier("")
	if err == nil {
		t.Error("error should be raised when email is empty")
	}
}
