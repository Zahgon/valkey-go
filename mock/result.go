package mock

import (
	"bufio"
	"bytes"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkey-io/valkey-go"
)

func Result(val valkey.ValkeyMessage) valkey.ValkeyResult {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func ErrorResult(err error) valkey.ValkeyResult {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResult)
}

func ValkeyString(v string) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyBlobString(v string) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyError(v string) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyInt64(v int64) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyFloat64(v float64) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyBool(v bool) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyNil() valkey.ValkeyMessage { _ = "STUB: not implemented"; return *new(valkey.ValkeyMessage) }

func ValkeyArray(values ...valkey.ValkeyMessage) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func ValkeyMap(kv map[string]valkey.ValkeyMessage) valkey.ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyMessage)
}

func serialize(m message, buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func ValkeyResultStreamError(err error) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

func ValkeyResultStream(ms ...valkey.ValkeyMessage) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

func MultiValkeyResultStream(ms ...valkey.ValkeyMessage) valkey.MultiValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.MultiValkeyResultStream)
}

func MultiValkeyResultStreamError(err error) valkey.ValkeyResultStream {
	_ = "STUB: not implemented"
	return *new(valkey.ValkeyResultStream)
}

type message struct {
	attrs   *valkey.ValkeyMessage
	bytes   *byte
	array   *valkey.ValkeyMessage
	integer int64
	typ     byte
	ttl     [7]byte
}

func (m *message) string() string { _ = "STUB: not implemented"; return "" }

func (m *message) values() []valkey.ValkeyMessage { _ = "STUB: not implemented"; return nil }

func slicemsg(typ byte, values []valkey.ValkeyMessage) message {
	_ = "STUB: not implemented"
	return *new(message)
}

func strmsg(typ byte, value string) message { _ = "STUB: not implemented"; return *new(message) }

type result struct {
	err error
	val valkey.ValkeyMessage
}

type pool struct {
	dead    any
	cond    *sync.Cond
	timer   *time.Timer
	make    func() any
	list    []any
	cleanup time.Duration
	size    int
	minSize int
	cap     int
	down    bool
	timerOn bool
}

type pipe struct {
	conn            net.Conn
	clhks           atomic.Value // closed hook, invoked after the conn is closed
	queue           any
	cache           any
	pshks           atomic.Pointer[pshks] // pubsub hook, registered by the SetPubSubHooks
	error           atomic.Pointer[errs]
	r               *bufio.Reader
	w               *bufio.Writer
	close           chan struct{}
	onInvalidations func([]valkey.ValkeyMessage)
	ssubs           *any // pubsub smessage subscriptions
	nsubs           *any // pubsub  message subscriptions
	psubs           *any // pubsub pmessage subscriptions
	r2p             *any
	pingTimer       *time.Timer // timer for background ping
	lftmTimer       *time.Timer // lifetime timer
	info            map[string]valkey.ValkeyMessage
	timeout         time.Duration
	pinggap         time.Duration
	maxFlushDelay   time.Duration
	lftm            time.Duration // lifetime
	wrCounter       atomic.Uint64
	version         int32
	blcksig         int32
	state           int32
	bgState         int32
	r2ps            bool // identify this pipe is used for resp2 pubsub or not
	noNoDelay       bool
	optIn           bool
}

type stream struct {
	p *pool
	w *pipe
	e error
	n int
}

type errs struct{ error }

type pshks struct {
	hooks valkey.PubSubHooks
	close chan error
}
