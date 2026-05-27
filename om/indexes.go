package om

import (
	"context"

	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/internal/cmds"
)

// createAndAliasIndex creates a new versioned index, aliases it to idx, and then drops all
// existing versioned indexes (idx_vN) that were previously associated with that alias.
func createAndAliasIndex(ctx context.Context, idx string, client valkey.Client, createCmd func(idx string) cmds.FtCreatePrefixPrefix, cmdFn func(schema FtCreateSchema) valkey.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// Create the new index
