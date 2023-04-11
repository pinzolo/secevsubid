package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestAccountIdentifier(t *testing.T) {
	wantUri := "acct:example.user@service.example.com"
	id, err := secevsubid.NewAccountIdentifier(wantUri)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatAccount {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatAccount)
	}
	if id.Uri() != wantUri {
		t.Errorf("invalid uri: got = %s, want = %s", id.Uri(), wantUri)
	}

	wantJSON := fmt.Sprintf(`{"format":"account","uri":"%s"}`, wantUri)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestAccountIdentifierWithEmptyUri(t *testing.T) {
	_, err := secevsubid.NewAccountIdentifier("")
	if err == nil {
		t.Error("error should be raised when uri is empty")
	}
}
