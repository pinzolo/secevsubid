package secevsubid_test

import (
	"encoding/json"
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestPhoneNumberIdentifier(t *testing.T) {
	wantPhoneNumber := "+12065550100"
	id, err := secevsubid.NewPhoneNumberIdentifier(wantPhoneNumber)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatPhoneNumber {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatPhoneNumber)
	}
	if id.PhoneNumber() != wantPhoneNumber {
		t.Errorf("invalid phoneNumber: got = %s, want = %s", id.PhoneNumber(), wantPhoneNumber)
	}

	wantJSON := fmt.Sprintf(`{"format":"phone_number","phone_number":"%s"}`, wantPhoneNumber)
	b, err := json.Marshal(id)
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestPhoneNumberIdentifierWithEmptyPhoneNumber(t *testing.T) {
	_, err := secevsubid.NewPhoneNumberIdentifier("")
	if err == nil {
		t.Error("error should be raised when phone number is empty")
	}
}
