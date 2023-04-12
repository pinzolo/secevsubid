package secevsubid_test

import (
	"encoding/json"
	"github.com/pinzolo/secevsubid"
	"reflect"
	"testing"
)

func TestNewAliasesIdentifier(t *testing.T) {
	account, _ := secevsubid.NewAccountIdentifier("acct:example.user@service.example.com")
	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	opaque, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	opaque2, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	aliases, _ := secevsubid.NewAliasesIdentifier()

	tests := []struct {
		name    string
		ids     []secevsubid.SubjectIdentifier
		wantErr bool
	}{
		{
			name:    "without identifiers",
			ids:     []secevsubid.SubjectIdentifier{},
			wantErr: false,
		},
		{
			name:    "with identifiers",
			ids:     []secevsubid.SubjectIdentifier{account, email, opaque},
			wantErr: false,
		},
		{
			name:    "nested aliases",
			ids:     []secevsubid.SubjectIdentifier{account, email, opaque, aliases},
			wantErr: true,
		},
		{
			name:    "duplicated identifier",
			ids:     []secevsubid.SubjectIdentifier{account, email, opaque, opaque2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := secevsubid.NewAliasesIdentifier(tt.ids...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAliasesIdentifier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Format() != secevsubid.FormatAliases {
					t.Errorf("invalid format: got = %s, want = %s", got.Format(), secevsubid.FormatAliases)
				}
				if !reflect.DeepEqual(got.Identifiers(), tt.ids) {
					t.Errorf("NewAliasesIdentifier() got = %v, want %v", got, tt.ids)
				}
			}
		})
	}
}

func TestAliasesIdentifier_Validate(t *testing.T) {
	id, _ := secevsubid.NewAliasesIdentifier()
	if err := id.Validate(); err == nil {
		t.Error("error should be raise when identifiers is empty")
	}

	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	err := id.AddIdentifier(email)
	if err != nil {
		t.Error(err)
	}
	if err = id.Validate(); err != nil {
		t.Error(err)
	}
}

func Test_aliasesIdentifier_ContainsIdentifier(t *testing.T) {
	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	opaque, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	opaque2, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	id, _ := secevsubid.NewAliasesIdentifier(opaque)

	tests := []struct {
		name string
		arg  secevsubid.SubjectIdentifier
		want bool
	}{
		{
			name: "not exist",
			arg:  email,
			want: false,
		},
		{
			name: "exists",
			arg:  opaque2,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := id.ContainsIdentifier(tt.arg); got != tt.want {
				t.Errorf("ContainsIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAliasesIdentifier_AddIdentifier(t *testing.T) {
	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	opaque, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	opaque2, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")
	tests := []struct {
		name    string
		arg     secevsubid.SubjectIdentifier
		wantIds []secevsubid.SubjectIdentifier
		wantErr bool
	}{
		{
			name:    "not exist",
			arg:     email,
			wantIds: []secevsubid.SubjectIdentifier{opaque, email},
			wantErr: false,
		},
		{
			name:    "exists",
			arg:     opaque2,
			wantIds: []secevsubid.SubjectIdentifier{opaque},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, _ := secevsubid.NewAliasesIdentifier(opaque)
			if err := id.AddIdentifier(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("AddIdentifier() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := id.Identifiers()
			if !reflect.DeepEqual(got, tt.wantIds) {
				t.Errorf("Identifers() got = %v, want %v", got, tt.wantIds)
			}
		})
	}
}

func TestAliasesIdentifier_MarshalJSON(t *testing.T) {
	account, _ := secevsubid.NewAccountIdentifier("acct:example.user@service.example.com")
	email, _ := secevsubid.NewEmailIdentifier("user@example.com")
	opaque, _ := secevsubid.NewOpaqueIdentifier("11112222333344445555")

	id, err := secevsubid.NewAliasesIdentifier(account, email, opaque)
	if err != nil {
		t.Error(err)
		return
	}
	b, err := json.Marshal(id)
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
		return
	}
	got := string(b)
	want := `{"format":"aliases","identifiers":[{"format":"account","uri":"acct:example.user@service.example.com"},{"format":"email","email":"user@example.com"},{"format":"opaque","id":"11112222333344445555"}]}`
	if got != want {
		t.Errorf("MarshalJSON() got = %v, want %v", got, want)
	}
}
