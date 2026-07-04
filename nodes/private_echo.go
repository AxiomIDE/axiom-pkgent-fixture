package nodes

import (
	"context"

	"pkgent001owner/entitlement-fixture/axiom"
	gen "pkgent001owner/entitlement-fixture/gen"
)

// PrivateEcho echoes the input text back unchanged. It exists only so the
// ENT-E2E-CROSSTENANT-UX e2e spec has a real, privately-owned node ULID to
// probe the package-entitlement compile gate against — its behavior is not
// under test.
func PrivateEcho(ctx context.Context, ax axiom.Context, input *gen.EntFixtureMessage) (*gen.EntFixtureMessage, error) {
	return &gen.EntFixtureMessage{Text: input.GetText()}, nil
}
