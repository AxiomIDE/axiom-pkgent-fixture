package nodes

import (
	"context"
	"testing"

	"pkgent001owner/entitlement-fixture/axiom"
	gen "pkgent001owner/entitlement-fixture/gen"
)

func TestPrivateEcho(t *testing.T) {
	out, err := PrivateEcho(context.Background(), axiom.Context{}, &gen.EntFixtureMessage{Text: "hello"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.GetText() != "hello" {
		t.Fatalf("expected text %q, got %q", "hello", out.GetText())
	}
}
