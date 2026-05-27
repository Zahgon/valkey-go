package valkeyhook

import (
	"context"
	"time"

	"github.com/valkey-io/valkey-go"
)

var _ valkey.Client = (*hookclient)(nil)

// Hook allows user to intercept valkey.Client by using WithHook
type Hook interface {
	Do(client valkey.Client, ctx context.Context, cmd valkey.Completed) (resp valkey.ValkeyResult)
	DoMulti(client valkey.Client, ctx context.Context, multi ...valkey.Completed) (resps []valkey.ValkeyResult)
	DoCache(client valkey.Client, ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) (resp valkey.ValkeyResult)
	DoMultiCache(client valkey.Client, ctx context.Context, multi ...valkey.CacheableTTL) (resps []valkey.ValkeyResult)
	Receive(client valkey.Client, ctx context.Context, subscribe valkey.Completed, fn func(msg valkey.PubSubMessage)) (err error)
	DoStream(client valkey.Client, ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream
	DoMultiStream(client valkey.Client, ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream
}

// WithHook wraps valkey.Client with Hook and allows the user to intercept valkey.Client
func WithHook(client valkey.Client, hook Hook) valkey.Client {
	_ = "STUB: not implemented"
	return *new(valkey.Client)
}

type hookclient struct {
	client valkey.Client
	hook   Hook
}

func (c *hookclient) B() valkey.Builder { _ = "STUB: not implemented"; return *new(valkey.Builder) }

func (c *hookclient) Do(ctx context.Context, cmd valkey.Completed) (resp valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func (c *hookclient) DoMulti(ctx context.Context, multi ...valkey.Completed) (resp []valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) DoCache(ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) (resp valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func (c *hookclient) DoMultiCache(ctx context.Context, multi ...valkey.CacheableTTL) (resps []valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) DoStream(ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

func (c *hookclient) DoMultiStream(ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.MultiValkeyResultStream)
}

func (c *hookclient) Dedicated(fn func(valkey.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) Dedicate() (valkey.DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(valkey.DedicatedClient), nil
}

func (c *hookclient) Receive(ctx context.Context, subscribe valkey.Completed, fn func(msg valkey.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) Nodes() map[string]valkey.Client { _ = "STUB: not implemented"; return nil }

func (c *hookclient) Mode() valkey.ClientMode {
	_ = "STUB: not implemented"
	return *new(valkey.ClientMode)
}

func (c *hookclient) Close() { _ = "STUB: not implemented"; return }

var _ valkey.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client *extended
	hook   Hook
}

func (d *dedicated) B() valkey.Builder { _ = "STUB: not implemented"; return *new(valkey.Builder) }

func (d *dedicated) Do(ctx context.Context, cmd valkey.Completed) (resp valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...valkey.Completed) (resp []valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Receive(ctx context.Context, subscribe valkey.Completed, fn func(msg valkey.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetPubSubHooks(hooks valkey.PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetOnInvalidations(fn func([]valkey.ValkeyMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Close() { _ = "STUB: not implemented"; return }

var _ valkey.Client = (*extended)(nil)

type extended struct {
	valkey.DedicatedClient
}

func (e *extended) DoCache(ctx context.Context, cmd valkey.Cacheable, ttl time.Duration) (resp valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func (e *extended) DoMultiCache(ctx context.Context, multi ...valkey.CacheableTTL) (resp []valkey.ValkeyResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *extended) DoStream(ctx context.Context, cmd valkey.Completed) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

func (c *extended) DoMultiStream(ctx context.Context, multi ...valkey.Completed) valkey.MultiValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.MultiValkeyResultStream)
}

func (e *extended) Dedicated(fn func(valkey.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *extended) Dedicate() (client valkey.DedicatedClient, cancel func()) {
	_ = "STUB: not implemented"
	return *new(valkey.DedicatedClient), nil
}

func (e *extended) Nodes() map[string]valkey.Client { _ = "STUB: not implemented"; return nil }

func (e *extended) Mode() valkey.ClientMode {
	_ = "STUB: not implemented"
	return *new(valkey.ClientMode)
}

type result struct {
	err error
	val valkey.ValkeyMessage
}

func NewErrorResult(err error) valkey.ValkeyResult {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

type stream struct {
	p *int
	w *int
	e error
	n int
}

func NewErrorResultStream(err error) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}
