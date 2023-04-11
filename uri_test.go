package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestUriIdentifier(t *testing.T) {
	wantUri := "acct:example.user@service.example.com"
	id, err := secevsubid.NewUriIdentifier(wantUri)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatUri {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatUri)
	}
	if id.Uri() != wantUri {
		t.Errorf("invalid uri: got = %s, want = %s", id.Uri(), wantUri)
	}

	wantJSON := fmt.Sprintf(`{"format":"uri","uri":"%s"}`, wantUri)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestUriIdentifierWithEmptyUri(t *testing.T) {
	_, err := secevsubid.NewUriIdentifier("")
	if err == nil {
		t.Error("error should be raised when uri is empty")
	}
}
