package valkey

import (
	"bytes"
	"errors"
	"io"
	"time"
	"unsafe"
)

const messageStructSize = int(unsafe.Sizeof(ValkeyMessage{}))

// Nil represents a Valkey Nil message
var Nil = &ValkeyError{typ: typeNull}

// ErrParse is a parse error that occurs when a Valkey message cannot be parsed correctly.
var errParse = errors.New("valkey: parse error")

// IsValkeyNil is a handy method to check if the error is a valkey nil response.
// All valkey nil responses returned as an error.
func IsValkeyNil(err error) bool {
	_ = "STUB: not implemented"

	// IsParseErr checks if the error is a parse error
	return false
}

func IsParseErr(err error) bool { _ = "STUB: not implemented"; return false }

// IsValkeyBusyGroup checks if it is a valkey BUSYGROUP message.
func IsValkeyBusyGroup(err error) bool { _ = "STUB: not implemented"; return false }

// IsValkeyErr is a handy method to check if the error is a valkey ERR response.
func IsValkeyErr(err error) (ret *ValkeyError, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ValkeyError is an error response or a nil message from the valkey instance
type ValkeyError ValkeyMessage

// string retrieves the contained string of the ValkeyError
func (m *ValkeyError) string() string { _ = "STUB: not implemented"; return "" }

func (r *ValkeyError) Error() string { _ = "STUB: not implemented"; return "" }

// IsNil checks if it is a valkey nil message.
func (r *ValkeyError) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsMoved checks if it is a valkey MOVED message and returns the moved address.
func (r *ValkeyError) IsMoved() (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

// IsAsk checks if it is a valkey ASK message and returns ask address.
func (r *ValkeyError) IsAsk() (addr string, ok bool) { _ = "STUB: not implemented"; return "", false }

// IsRedirect checks if it is a valkey REDIRECT message and returns redirect address.
func (r *ValkeyError) IsRedirect() (addr string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func fixIPv6HostPort(addr string) string { _ = "STUB: not implemented"; return "" }

// skip ipv4 and enclosed ipv6

// IsTryAgain checks if it is a valkey TRYAGAIN message and returns ask address.
func (r *ValkeyError) IsTryAgain() bool { _ = "STUB: not implemented"; return false }

// IsLoading checks if it is a valkey LOADING message
func (r *ValkeyError) IsLoading() bool { _ = "STUB: not implemented"; return false }

// IsClusterDown checks if it is a valkey CLUSTERDOWN message and returns ask address.
func (r *ValkeyError) IsClusterDown() bool { _ = "STUB: not implemented"; return false }

// IsNoScript checks if it is a valkey NOSCRIPT message.
func (r *ValkeyError) IsNoScript() bool { _ = "STUB: not implemented"; return false }

// IsBusyGroup checks if it is a valkey BUSYGROUP message.
func (r *ValkeyError) IsBusyGroup() bool { _ = "STUB: not implemented"; return false }

func newResult(val ValkeyMessage, err error) ValkeyResult {
	_ = "STUB: not implemented"
	return *new(ValkeyResult)
}

func newErrResult(err error) ValkeyResult { _ = "STUB: not implemented"; return *new(ValkeyResult) }

// ValkeyResult is the return struct from Client.Do or Client.DoCache
// it contains either a valkey response or an underlying error (ex. network timeout).
type ValkeyResult struct {
	err error
	val ValkeyMessage
}

// NonValkeyError can be used to check if there is an underlying error (ex. network timeout).
func (r ValkeyResult) NonValkeyError() error {
	_ = "STUB: not implemented"

	// Error returns either underlying error or valkey error or nil
	return nil
}

func (r ValkeyResult) Error() (err error) { _ = "STUB: not implemented"; return nil }

// ToMessage retrieves the ValkeyMessage
func (r ValkeyResult) ToMessage() (v ValkeyMessage, err error) {
	_ = "STUB: not implemented"
	return *new(ValkeyMessage), nil
}

// ToInt64 delegates to ValkeyMessage.ToInt64
func (r ValkeyResult) ToInt64() (v int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToBool delegates to ValkeyMessage.ToBool
func (r ValkeyResult) ToBool() (v bool, err error) { _ = "STUB: not implemented"; return false, nil }

// ToFloat64 delegates to ValkeyMessage.ToFloat64
func (r ValkeyResult) ToFloat64() (v float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToString delegates to ValkeyMessage.ToString
func (r ValkeyResult) ToString() (v string, err error) { _ = "STUB: not implemented"; return "", nil }

// AsReader delegates to ValkeyMessage.AsReader
func (r ValkeyResult) AsReader() (v io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// AsBytes delegates to ValkeyMessage.AsBytes
func (r ValkeyResult) AsBytes() (v []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// DecodeJSON delegates to ValkeyMessage.DecodeJSON
func (r ValkeyResult) DecodeJSON(v any) (err error) { _ = "STUB: not implemented"; return nil }

// AsInt64 delegates to ValkeyMessage.AsInt64
func (r ValkeyResult) AsInt64() (v int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsUint64 delegates to ValkeyMessage.AsUint64
func (r ValkeyResult) AsUint64() (v uint64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsBool delegates to ValkeyMessage.AsBool
func (r ValkeyResult) AsBool() (v bool, err error) { _ = "STUB: not implemented"; return false, nil }

// AsFloat64 delegates to ValkeyMessage.AsFloat64
func (r ValkeyResult) AsFloat64() (v float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToArray delegates to ValkeyMessage.ToArray
func (r ValkeyResult) ToArray() (v []ValkeyMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrSlice delegates to ValkeyMessage.AsStrSlice
func (r ValkeyResult) AsStrSlice() (v []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntSlice delegates to ValkeyMessage.AsIntSlice
func (r ValkeyResult) AsIntSlice() (v []int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsFloatSlice delegates to ValkeyMessage.AsFloatSlice
func (r ValkeyResult) AsFloatSlice() (v []float64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsBoolSlice delegates to ValkeyMessage.AsBoolSlice
func (r ValkeyResult) AsBoolSlice() (v []bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRangeEntry delegates to ValkeyMessage.AsXRangeEntry
func (r ValkeyResult) AsXRangeEntry() (v XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return *new(XRangeEntry), nil
}

// AsXRange delegates to ValkeyMessage.AsXRange
func (r ValkeyResult) AsXRange() (v []XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsZScore delegates to ValkeyMessage.AsZScore
func (r ValkeyResult) AsZScore() (v ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScores delegates to ValkeyMessage.AsZScores
func (r ValkeyResult) AsZScores() (v []ZScore, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRead delegates to ValkeyMessage.AsXRead
func (r ValkeyResult) AsXRead() (v map[string][]XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRangeSlice delegates to ValkeyMessage.AsXRangeSlice
func (r ValkeyResult) AsXRangeSlice() (v XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return *new(XRangeSlice), nil
}

// AsXRangeSlices delegates to ValkeyMessage.AsXRangeSlices
func (r ValkeyResult) AsXRangeSlices() (v []XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXReadSlices delegates to ValkeyMessage.AsXReadSlices
func (r ValkeyResult) AsXReadSlices() (v map[string][]XRangeSlice, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ValkeyResult) AsLMPop() (v KeyValues, err error) {
	_ = "STUB: not implemented"
	return *new(KeyValues), nil
}

func (r ValkeyResult) AsZMPop() (v KeyZScores, err error) {
	_ = "STUB: not implemented"
	return *new(KeyZScores), nil
}

func (r ValkeyResult) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (r ValkeyResult) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (r ValkeyResult) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

func (r ValkeyResult) AsGeosearch() (locations []GeoLocation, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsMap delegates to ValkeyMessage.AsMap
func (r ValkeyResult) AsMap() (v map[string]ValkeyMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrMap delegates to ValkeyMessage.AsStrMap
func (r ValkeyResult) AsStrMap() (v map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntMap delegates to ValkeyMessage.AsIntMap
func (r ValkeyResult) AsIntMap() (v map[string]int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsScanEntry delegates to ValkeyMessage.AsScanEntry.
func (r ValkeyResult) AsScanEntry() (v ScanEntry, err error) {
	_ = "STUB: not implemented"
	return *new(ScanEntry), nil
}

// ToMap delegates to ValkeyMessage.ToMap
func (r ValkeyResult) ToMap() (v map[string]ValkeyMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToAny delegates to ValkeyMessage.ToAny
func (r ValkeyResult) ToAny() (v any, err error) { _ = "STUB: not implemented"; return *new(any), nil }

// IsCacheHit delegates to ValkeyMessage.IsCacheHit
func (r ValkeyResult) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

// CacheTTL delegates to ValkeyMessage.CacheTTL
func (r ValkeyResult) CacheTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePTTL delegates to ValkeyMessage.CachePTTL
func (r ValkeyResult) CachePTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePXAT delegates to ValkeyMessage.CachePXAT
func (r ValkeyResult) CachePXAT() int64 { _ = "STUB: not implemented"; return 0 }

// String returns human-readable representation of ValkeyResult
func (r *ValkeyResult) String() string { _ = "STUB: not implemented"; return "" }

type prettyValkeyResult ValkeyResult

// MarshalJSON implements json.Marshaler interface
func (r *prettyValkeyResult) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValkeyMessage is a valkey response message, it may be a nil response
type ValkeyMessage struct {
	attrs *ValkeyMessage
	bytes *byte
	array *ValkeyMessage

	// intlen is used for a simple number or
	// in conjunction with an array or bytes to store the length of array or string
	intlen int64
	typ    byte
	ttl    [7]byte
}

func (m *ValkeyMessage) string() string { _ = "STUB: not implemented"; return "" }

func (m *ValkeyMessage) values() []ValkeyMessage { _ = "STUB: not implemented"; return nil }

func (m *ValkeyMessage) setString(s string) { _ = "STUB: not implemented"; return }

func (m *ValkeyMessage) setValues(values []ValkeyMessage) { _ = "STUB: not implemented"; return }

func (m *ValkeyMessage) cachesize() int {
	_ = "STUB: not implemented"
	// typ (1) + length (8) TODO: can we use VarInt instead of fixed 8 bytes for length?
	return 0
}

func (m *ValkeyMessage) serialize(o *bytes.Buffer) {
	_ = "STUB: not implemented"
	// TODO: can we use VarInt instead of fixed 8 bytes for length?
	return
}

var ErrCacheUnmarshal = errors.New("cache unmarshal error")

func (m *ValkeyMessage) unmarshalView(c int64, buf []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: can we use VarInt instead of fixed 8 bytes for length?

// CacheSize returns the buffer size needed by the CacheMarshal.
func (m *ValkeyMessage) CacheSize() int { _ = "STUB: not implemented"; return 0 }

// 7 for ttl

// CacheMarshal writes serialized ValkeyMessage to the provided buffer.
// If the provided buffer is nil, CacheMarshal will allocate one.
// Note that an output format is not compatible with different client versions.
func (m *ValkeyMessage) CacheMarshal(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

// CacheUnmarshalView construct the ValkeyMessage from the buffer produced by CacheMarshal.
// Note that the buffer can't be reused after CacheUnmarshalView since it uses unsafe.String on top of the buffer.
func (m *ValkeyMessage) CacheUnmarshalView(buf []byte) error { _ = "STUB: not implemented"; return nil }

// IsNil check if the message is a valkey nil response
func (m *ValkeyMessage) IsNil() bool { _ = "STUB: not implemented"; return false }

// IsInt64 check if the message is a valkey RESP3 int response
func (m *ValkeyMessage) IsInt64() bool { _ = "STUB: not implemented"; return false }

// IsFloat64 check if the message is a valkey RESP3 double response
func (m *ValkeyMessage) IsFloat64() bool { _ = "STUB: not implemented"; return false }

// IsString check if the message is a valkey string response
func (m *ValkeyMessage) IsString() bool { _ = "STUB: not implemented"; return false }

// IsBool check if the message is a valkey RESP3 bool response
func (m *ValkeyMessage) IsBool() bool { _ = "STUB: not implemented"; return false }

// IsArray check if the message is a valkey array response
func (m *ValkeyMessage) IsArray() bool { _ = "STUB: not implemented"; return false }

// IsMap check if the message is a valkey RESP3 map response
func (m *ValkeyMessage) IsMap() bool { _ = "STUB: not implemented"; return false }

// Error check if the message is a valkey error response, including nil response
func (m *ValkeyMessage) Error() error { _ = "STUB: not implemented"; return nil }

// kvrocks: https://github.com/redis/rueidis/issues/152#issuecomment-1333923750

// ToString check if the message is a valkey string response and return it
func (m *ValkeyMessage) ToString() (val string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AsReader check if the message is a valkey string response and wrap it with the strings.NewReader
func (m *ValkeyMessage) AsReader() (reader io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// AsBytes check if the message is a valkey string response and return it as an immutable []byte
func (m *ValkeyMessage) AsBytes() (bs []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeJSON check if the message is a valkey string response and treat it as JSON, then unmarshal it into the provided value
func (m *ValkeyMessage) DecodeJSON(v any) (err error) { _ = "STUB: not implemented"; return nil }

// AsInt64 check if the message is a valkey string response and parse it as int64
func (m *ValkeyMessage) AsInt64() (val int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// AsUint64 check if the message is a valkey string response and parse it as uint64
func (m *ValkeyMessage) AsUint64() (val uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AsBool checks if the message is a non-nil response and parses it as bool
func (m *ValkeyMessage) AsBool() (val bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AsFloat64 check if the message is a valkey string response and parse it as float64
func (m *ValkeyMessage) AsFloat64() (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToInt64 check if the message is a valkey RESP3 int response and return it
func (m *ValkeyMessage) ToInt64() (val int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ToBool check if the message is a valkey RESP3 bool response and return it
func (m *ValkeyMessage) ToBool() (val bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ToFloat64 check if the message is a valkey RESP3 double response and return it
func (m *ValkeyMessage) ToFloat64() (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToArray check if the message is a valkey array/set response and return it
func (m *ValkeyMessage) ToArray() ([]ValkeyMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrSlice check if the message is a valkey array/set response and convert to []string.
// valkey nil element and other non-string elements will be present as zero.
func (m *ValkeyMessage) AsStrSlice() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// AsIntSlice check if the message is a valkey array/set response and convert to []int64.
// valkey nil element and other non-integer elements will be present as zero.
func (m *ValkeyMessage) AsIntSlice() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

// AsFloatSlice check if the message is a valkey array/set response and convert to []float64.
// valkey nil element and other non-float elements will be present as zero.
func (m *ValkeyMessage) AsFloatSlice() ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsBoolSlice checks if the message is a valkey array/set response and converts it to []bool.
// Valkey nil elements and other non-boolean elements will be represented as false.
func (m *ValkeyMessage) AsBoolSlice() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// Ignore error, non-boolean values will be false

// XRangeEntry is the element type of both XRANGE and XREVRANGE command response array
type XRangeEntry struct {
	FieldValues map[string]string
	ID          string
}

// AsXRangeEntry check if the message is a valkey array/set response of length 2 and convert to XRangeEntry
func (m *ValkeyMessage) AsXRangeEntry() (XRangeEntry, error) {
	_ = "STUB: not implemented"
	return *new(XRangeEntry), nil
}

// AsXRange check if the message is a valkey array/set response and convert to []XRangeEntry
func (m *ValkeyMessage) AsXRange() ([]XRangeEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXRead converts XREAD/XREADGRUOP response to map[string][]XRangeEntry
func (m *ValkeyMessage) AsXRead() (ret map[string][]XRangeEntry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// New slice-based structures that preserve order and duplicates
type XRangeSlice struct {
	ID          string
	FieldValues []XRangeFieldValue
}

type XRangeFieldValue struct {
	Field string
	Value string
}

// AsXRangeSlice converts a ValkeyMessage to XRangeSlice (preserves order and duplicates)
func (m *ValkeyMessage) AsXRangeSlice() (XRangeSlice, error) {
	_ = "STUB: not implemented"
	return *new(XRangeSlice), nil
}

// Handle the field-values array

// Convert pairs to slice (preserving order)

// AsXRangeSlices converts multiple XRange entries to slice format
func (m *ValkeyMessage) AsXRangeSlices() ([]XRangeSlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsXReadSlices converts XREAD/XREADGROUP response to use slice format
func (m *ValkeyMessage) AsXReadSlices() (map[string][]XRangeSlice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ZScore is the element type of ZRANGE WITHSCORES, ZDIFF WITHSCORES and ZPOPMAX command response
type ZScore struct {
	Member string
	Score  float64
}

func toZScore(values []ValkeyMessage) (s ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScore converts ZPOPMAX and ZPOPMIN command with count 1 response to a single ZScore
func (m *ValkeyMessage) AsZScore() (s ZScore, err error) {
	_ = "STUB: not implemented"
	return *new(ZScore), nil
}

// AsZScores converts ZRANGE WITHSCORES, ZDIFF WITHSCORES and ZPOPMAX/ZPOPMIN command with count > 1 responses to []ZScore
func (m *ValkeyMessage) AsZScores() ([]ZScore, error) { _ = "STUB: not implemented"; return nil, nil }

// ScanEntry is the element type of both SCAN, SSCAN, HSCAN and ZSCAN command response.
type ScanEntry struct {
	Elements []string
	Cursor   uint64
}

// AsScanEntry check if the message is a valkey array/set response of length 2 and convert to ScanEntry.
func (m *ValkeyMessage) AsScanEntry() (e ScanEntry, err error) {
	_ = "STUB: not implemented"
	return *new(ScanEntry), nil
}

// AsMap check if the message is a valkey array/set response and convert to map[string]ValkeyMessage
func (m *ValkeyMessage) AsMap() (map[string]ValkeyMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsStrMap check if the message is a valkey map/array/set response and convert to map[string]string.
// valkey nil element and other non-string elements will be present as zero.
func (m *ValkeyMessage) AsStrMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsIntMap check if the message is a valkey map/array/set response and convert to map[string]int64.
// valkey nil element and other non-integer elements will be present as zero.
func (m *ValkeyMessage) AsIntMap() (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type KeyValues struct {
	Key    string
	Values []string
}

func (m *ValkeyMessage) AsLMPop() (kvs KeyValues, err error) {
	_ = "STUB: not implemented"
	return *new(KeyValues), nil
}

type KeyZScores struct {
	Key    string
	Values []ZScore
}

func (m *ValkeyMessage) AsZMPop() (kvs KeyZScores, err error) {
	_ = "STUB: not implemented"
	return *new(KeyZScores), nil
}

type FtSearchDoc struct {
	Doc   map[string]string
	Key   string
	Score float64
}

func (m *ValkeyMessage) AsFtSearch() (total int64, docs []FtSearchDoc, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (m *ValkeyMessage) AsFtAggregate() (total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (m *ValkeyMessage) AsFtAggregateCursor() (cursor, total int64, docs []map[string]string, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil, nil
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

func (m *ValkeyMessage) AsGeosearch() ([]GeoLocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//name

//distance

//hash

//coordinates

// ToMap check if the message is a valkey RESP3 map response and return it
func (m *ValkeyMessage) ToMap() (map[string]ValkeyMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToAny turns the message into go any value
func (m *ValkeyMessage) ToAny() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// IsCacheHit check if the message is from the client side cache
func (m *ValkeyMessage) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

// CacheTTL returns the remaining TTL in seconds of client side cache
func (m *ValkeyMessage) CacheTTL() (ttl int64) { _ = "STUB: not implemented"; return 0 }

// CachePTTL returns the remaining PTTL in seconds of client side cache
func (m *ValkeyMessage) CachePTTL() int64 { _ = "STUB: not implemented"; return 0 }

// CachePXAT returns the remaining PXAT in seconds of client side cache
func (m *ValkeyMessage) CachePXAT() int64 { _ = "STUB: not implemented"; return 0 }

func (m *ValkeyMessage) relativePTTL(now time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func (m *ValkeyMessage) getExpireAt() int64 { _ = "STUB: not implemented"; return 0 }

func (m *ValkeyMessage) setExpireAt(pttl int64) { _ = "STUB: not implemented"; return }

func toMap(values []ValkeyMessage) (map[string]ValkeyMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ValkeyMessage) approximateSize() (s int) { _ = "STUB: not implemented"; return 0 }

// String returns the human-readable representation of ValkeyMessage
func (m *ValkeyMessage) String() string { _ = "STUB: not implemented"; return "" }

type prettyValkeyMessage ValkeyMessage

func (m *prettyValkeyMessage) string() string { _ = "STUB: not implemented"; return "" }

func (m *prettyValkeyMessage) values() []ValkeyMessage { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements json.Marshaler interface
func (m *prettyValkeyMessage) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func slicemsg(typ byte, values []ValkeyMessage) ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(ValkeyMessage)
}

func strmsg(typ byte, value string) ValkeyMessage {
	_ = "STUB: not implemented"
	return *new(ValkeyMessage)
}
