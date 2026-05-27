package valkeycompat

import (
	"context"
	"errors"
	"time"

	"github.com/valkey-io/valkey-go"
)

var TxFailedErr = errors.New("valkey: transaction failed")

var _ Pipeliner = (*TxPipeline)(nil)

type rePipeline = Pipeline

func newTxPipeline(real valkey.Client) *TxPipeline { _ = "STUB: not implemented"; return nil }

type TxPipeline struct {
	*rePipeline
}

func (c *TxPipeline) Exec(ctx context.Context) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *TxPipeline) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

var _ valkey.Client = (*txproxy)(nil)

type txproxy struct {
	valkey.CoreClient
}

func (p *txproxy) DoCache(_ context.Context, _ valkey.Cacheable, _ time.Duration) (resp valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func (p *txproxy) DoMultiCache(_ context.Context, _ ...valkey.CacheableTTL) (resp []valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return nil
}

func (p *txproxy) DoStream(_ context.Context, _ valkey.Completed) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

func (p *txproxy) DoMultiStream(_ context.Context, _ ...valkey.Completed) valkey.MultiValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.MultiValkeyResultStream)
}

func (p *txproxy) Dedicated(_ func(valkey.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *txproxy) Dedicate() (client valkey.DedicatedClient, cancel func()) {
	_ = "STUB: not implemented"
	return *new(valkey.DedicatedClient), nil
}

func (p *txproxy) Nodes() map[string]valkey.Client { _ = "STUB: not implemented"; return nil }

func (p *txproxy) Mode() valkey.ClientMode {
	_ = "STUB: not implemented"
	return *new(valkey.ClientMode)
}

type Tx interface {
	CoreCmdable
	Watch(ctx context.Context, keys ...string) *StatusCmd
	Unwatch(ctx context.Context, keys ...string) *StatusCmd
	Close(ctx context.Context) error
}

func newTx(client valkey.DedicatedClient, cancel func()) *tx { _ = "STUB: not implemented"; return nil }

type tx struct {
	CoreCmdable
	cancel func()
}

func (t *tx) Watch(ctx context.Context, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (t *tx) Unwatch(ctx context.Context, _ ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (t *tx) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }
