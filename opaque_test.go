package secevsubid_test

import (
	"fmt"
	"github.com/pinzolo/secevsubid"
	"testing"
)

func TestOpaqueIdentifier(t *testing.T) {
	wantId := "11112222333344445555"
	id, err := secevsubid.NewOpaqueIdentifier(wantId)
	if err != nil {
		t.Error(err)
		return
	}

	if id.Format() != secevsubid.FormatOpaque {
		t.Errorf("invalid format: got = %s, want = %s", id.Format(), secevsubid.FormatOpaque)
	}
	if id.Id() != wantId {
		t.Errorf("invalid id: got = %s, want = %s", id.Id(), wantId)
	}

	wantJSON := fmt.Sprintf(`{"format":"opaque","id":"%s"}`, wantId)
	b, err := id.MarshalJSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != wantJSON {
		t.Errorf("invalid JSON conversion: got = %s, want = %s", string(b), wantJSON)
	}
}

func TestOpaqueIdentifierWithEmptyId(t *testing.T) {
	_, err := secevsubid.NewOpaqueIdentifier("")
	if err == nil {
		t.Error("error should be raised when id is empty")
	}
}
