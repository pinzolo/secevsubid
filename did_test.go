package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestDidIdentifier(t *testing.T) {
	wantUrl := "did:example:123456"
	id, err := secevsubid.NewDidIdentifier(wantUrl)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatDid {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatDid)
	}
	if id.Url() != wantUrl {
		t.Errorf("invalid url: got = %s, want = %s", id.Url(), wantUrl)
	}

	wantJSON := fmt.Sprintf(`{"format":"did","url":"%s"}`, wantUrl)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestDidIdentifierWithEmptyUrl(t *testing.T) {
	_, err := secevsubid.NewDidIdentifier("")
	if err == nil {
		t.Error("error should be raised when url is empty")
	}
}
