package valkey

import (
	"context"
)

type queuedCmd struct {
	ch    chan ValkeyResult
	one   Completed
	multi []Completed
	resps []ValkeyResult
}

type flowBuffer struct {
	f chan queuedCmd
	r chan queuedCmd
	w chan queuedCmd
	c *chan ValkeyResult
}

var _ queue = (*flowBuffer)(nil)

func newFlowBuffer(factor int) *flowBuffer { _ = "STUB: not implemented"; return nil }

func (b *flowBuffer) PutOne(ctx context.Context, m Completed) (chan ValkeyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *flowBuffer) PutMulti(ctx context.Context, m []Completed, resps []ValkeyResult) (chan ValkeyResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextWriteCmd should be only called by one dedicated thread
func (b *flowBuffer) NextWriteCmd() (one Completed, multi []Completed, ch chan ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil
}

// WaitForWrite should be only called by one dedicated thread
func (b *flowBuffer) WaitForWrite() (one Completed, multi []Completed, ch chan ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil
}

// NextResultCh should be only called by one dedicated thread
func (b *flowBuffer) NextResultCh() (one Completed, multi []Completed, ch chan ValkeyResult, resps []ValkeyResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil, nil
}

// FinishResult should be only called by one dedicated thread
func (b *flowBuffer) FinishResult() { _ = "STUB: not implemented"; return }
