package secevsubid_test

import (
	"encoding/json"
	"github.com/pinzolo/secevsubid"
	"reflect"
	"testing"
)

func TestWrapper(t *testing.T) {
	id, _ := secevsubid.NewOpaqueIdentifier("id")
	w := secevsubid.NewWrapper(id)
	got := w.Value()
	if !reflect.DeepEqual(got, id) {
		t.Errorf("Value() = %v, want %v", got, id)
	}
}

func TestWrapper_MarshalJSON(t *testing.T) {
	id, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	w := secevsubid.NewWrapper(id)
	b, err := json.Marshal(w)
	if err != nil {
		t.Error(err)
		return
	}

	got := string(b)
	want := `{"format":"opaque","id":"11112222333344445555"}`
	if got != want {
		t.Errorf("MarshalJSON() = %v, want %v", got, want)
	}
}

func TestWrapper_UnmarshalJSON(t *testing.T) {
	account, _ := secevsubid.NewAccountIdentifier("acct:example.user@service.example.com")
	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	issSub, _ := secevsubid.NewIssuerSubjectIdentifier("https://issuer.example.com/", "145234573")
	opaque, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	phoneNum, _ := secevsubid.NewPhoneNumberIdentifier("+12065550100")
	did, _ := secevsubid.NewDidIdentifier("did:example:123456")
	uri, _ := secevsubid.NewUriIdentifier("https://user.example.com/")

	tests := []struct {
		name    string
		json    string
		want    secevsubid.SubjectIdentifier
		wantErr bool
	}{
		{
			name: "account success",
			json: `
{
  "format": "account",
  "uri": "acct:example.user@service.example.com"
}`,
			want:    account,
			wantErr: false,
		},
		{
			name: "email success",
			json: `
{
  "format": "email",
  "email": "user@example.com"
}`,
			want:    email,
			wantErr: false,
		},
		{
			name: "issuer and subject success",
			json: `
{
  "format": "iss_sub",
  "iss": "https://issuer.example.com/",
  "sub": "145234573"
}`,
			want:    issSub,
			wantErr: false,
		},
		{
			name: "opaque success",
			json: `
{
  "format": "opaque",
  "id": "11112222333344445555"
}`,
			want:    opaque,
			wantErr: false,
		},
		{
			name: "phone number success",
			json: `
{
  "format": "phone_number",
  "phone_number": "+12065550100"
}`,
			want:    phoneNum,
			wantErr: false,
		},
		{
			name: "did success",
			json: `
{
  "format": "did",
  "url": "did:example:123456"
}`,
			want:    did,
			wantErr: false,
		},
		{
			name: "uri success",
			json: `
{
  "format": "uri",
  "uri": "https://user.example.com/"
}`,
			want:    uri,
			wantErr: false,
		},
		{
			name: "no format field",
			json: `
{
  "uri": "https://user.example.com/"
}`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "no format value",
			json: `
{
  "format": "",
  "uri": "https://user.example.com/"
}`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "unknown format",
			json: `
{
  "format": "unknown",
  "uri": "https://user.example.com/"
}`,
			want:    nil,
			wantErr: true,
		},
		{
			name: "broken JSON",
			json: `
{
  "format": "uri",
  "uri": "https://user.example.com/"
`,
			want:    nil,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &secevsubid.Wrapper{}
			b := []byte(tt.json)
			if err := json.Unmarshal(b, &w); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
