package valkeyotel

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

type OpNameResolver interface {
	OpName(ctx context.Context, cmd valkey.Completed) string
	MultiOpName(ctx context.Context, cmds valkey.Commands) string
	MultiCacheableOpName(ctx context.Context, cmds []valkey.CacheableTTL) string
}

var _ OpNameResolver = (*DefaultOpNameResolver)(nil)

type DefaultOpNameResolver struct {
	// Limit controls how many elements are used to compose the resulting operation name.
	// If Limit is greater than zero, only the first Limit commands are used.
	Limit int
}

func (DefaultOpNameResolver) OpName(_ context.Context, cmd valkey.Completed) string {
	_ = "STUB: not implemented"
	return ""
}

func (r DefaultOpNameResolver) MultiOpName(_ context.Context, cmds valkey.Commands) string {
	_ = "STUB: not implemented"
	return ""
}

func (r DefaultOpNameResolver) MultiCacheableOpName(_ context.Context, cmds []valkey.CacheableTTL) string {
	_ = "STUB: not implemented"
	return ""
}
