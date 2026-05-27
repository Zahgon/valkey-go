package valkey

import (
	"context"
	"sync"
	"time"
)

// NewCacheStoreFn can be provided in ClientOption for using a custom CacheStore implementation
type NewCacheStoreFn func(CacheStoreOption) CacheStore

// CacheStoreOption will be passed to NewCacheStoreFn
type CacheStoreOption struct {
	// CacheSizeEachConn is valkey client side cache size that bind to each TCP connection to a single valkey instance.
	// The default is DefaultCacheBytes.
	CacheSizeEachConn int
}

// CacheStore is the store interface for the client side caching
// More detailed interface requirement can be found in cache_test.go
type CacheStore interface {
	// Flight is called when DoCache and DoMultiCache, with the requested client side ttl and the current time.
	// It should look up the store in a single-flight manner and return one of the following three combinations:
	// Case 1: (empty ValkeyMessage, nil CacheEntry)     <- when cache missed, and valkey will send the request to valkey.
	// Case 2: (empty ValkeyMessage, non-nil CacheEntry) <- when cache missed, and valkey will use CacheEntry.Wait to wait for response.
	// Case 3: (non-empty ValkeyMessage, nil CacheEntry) <- when cache hit
	Flight(key, cmd string, ttl time.Duration, now time.Time) (v ValkeyMessage, e CacheEntry)
	// Update is called when receiving the response of the request sent by the above Flight Case 1 from valkey.
	// It should not only update the store but also deliver the response to all CacheEntry.Wait and return a desired client side PXAT of the response.
	// Note that the server side expire time can be retrieved from ValkeyMessage.CachePXAT.
	Update(key, cmd string, val ValkeyMessage) (pxat int64)
	// Cancel is called when the request sent by the above Flight Case 1 failed.
	// It should not only deliver the error to all CacheEntry.Wait but also remove the CacheEntry from the store.
	Cancel(key, cmd string, err error)
	// Delete is called when receiving invalidation notifications from valkey.
	// If the keys are nil, then it should delete all non-pending cached entries under all keys.
	// If the keys are not nil, then it should delete all non-pending cached entries under those keys.
	Delete(keys []ValkeyMessage)
	// Close is called when the connection between valkey is broken.
	// It should flush all cached entries and deliver the error to all pending CacheEntry.Wait.
	Close(err error)
}

// CacheEntry should be used to wait for a single-flight response when cache missed.
type CacheEntry interface {
	Wait(ctx context.Context) (ValkeyMessage, error)
}

// SimpleCache is an alternative interface should be paired with NewSimpleCacheAdapter to construct a CacheStore
type SimpleCache interface {
	Get(key string) ValkeyMessage
	Set(key string, val ValkeyMessage)
	Del(key string)
	Flush()
}

// NewSimpleCacheAdapter converts a SimpleCache into CacheStore
func NewSimpleCacheAdapter(store SimpleCache) CacheStore {
	_ = "STUB: not implemented"
	return *new(CacheStore)
}

type adapter struct {
	store   SimpleCache
	flights map[string]map[string]CacheEntry
	mu      sync.RWMutex
}

func (a *adapter) Flight(key, cmd string, ttl time.Duration, now time.Time) (ValkeyMessage, CacheEntry) {
	_ = "STUB: not implemented"
	return *new(ValkeyMessage), *new(CacheEntry)
}

func (a *adapter) Update(key, cmd string, val ValkeyMessage) (sxat int64) {
	_ = "STUB: not implemented"
	return 0
}

func (a *adapter) Cancel(key, cmd string, err error) { _ = "STUB: not implemented"; return }

func (a *adapter) del(key string) { _ = "STUB: not implemented"; return }

func (a *adapter) Delete(keys []ValkeyMessage) { _ = "STUB: not implemented"; return }

func (a *adapter) Close(err error) { _ = "STUB: not implemented"; return }

type adapterEntry struct {
	err error
	ch  chan struct{}
	val ValkeyMessage
	xat int64
}

func (a *adapterEntry) set(val ValkeyMessage, err error) { _ = "STUB: not implemented"; return }

func (a *adapterEntry) Wait(ctx context.Context) (ValkeyMessage, error) {
	_ = "STUB: not implemented"
	return *new(ValkeyMessage), nil
}
