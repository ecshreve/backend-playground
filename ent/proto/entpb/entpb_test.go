package entpb_test

import (
	"testing"

	"github.com/ecshreve/backend-playground/ent/proto/entpb"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:  "eric",
		Email: "eric@example.com",
	}
	if user.GetName() != "eric" {
		t.Fatal("expected user name to be eric")
	}
	if user.GetEmail() != "eric@example.com" {
		t.Fatal("expected email address to be eric@example.com")
	}
}
