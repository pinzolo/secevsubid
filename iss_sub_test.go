package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestIssuerSubjectIdentifier(t *testing.T) {
	wantIss := "https://issuer.example.com/"
	wantSub := "145234573"
	id, err := secevsubid.NewIssuerSubjectIdentifier(wantIss, wantSub)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatIssuerSubject {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatIssuerSubject)
	}
	if id.Issuer() != wantIss {
		t.Errorf("invalid uri: got = %s, want = %s", id.Issuer(), wantIss)
	}
	if id.Subject() != wantSub {
		t.Errorf("invalid uri: got = %s, want = %s", id.Subject(), wantSub)
	}

	wantJSON := fmt.Sprintf(`{"format":"iss_sub","iss":"%s","sub":"%s"}`, wantIss, wantSub)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestIssuerSubjectIdentifierWithEmptyIssuer(t *testing.T) {
	_, err := secevsubid.NewIssuerSubjectIdentifier("", "145234573")
	if err == nil {
		t.Error("error should be raised when issuer is empty")
	}
}

func TestIssuerSubjectIdentifierWithEmptySubject(t *testing.T) {
	_, err := secevsubid.NewIssuerSubjectIdentifier("https://issuer.example.com/", "")
	if err == nil {
		t.Error("error should be raised when subject is empty")
	}
}
