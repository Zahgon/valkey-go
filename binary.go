package valkey

// BinaryString convert the provided []byte into a string without a copy. It does what strings.Builder.String() does.
// Valkey Strings are binary safe; this means that it is safe to store any []byte into Valkey directly.
// Users can use this BinaryString helper to insert a []byte as the part of valkey command. For example:
//
//	client.B().Set().Key(valkey.BinaryString([]byte{0})).Value(valkey.BinaryString([]byte{0})).Build()
//
// To read back the []byte of the string returned from the Valkey, it is recommended to use the ValkeyMessage.AsReader.
func BinaryString(bs []byte) string { _ = "STUB: not implemented"; return "" }

// VectorString32 convert the provided []float32 into a string. Users can use this to build vector search queries:
//
//	client.B().FtSearch().Index("idx").Query("*=>[KNN 5 @vec $V]").
//	    Params().Nargs(2).NameValue().NameValue("V", valkey.VectorString32([]float32{1})).
//	    Dialect(2).Build()
func VectorString32(v []float32) string { _ = "STUB: not implemented"; return "" }

// ToVector32 reverts VectorString32. User can use this to convert valkey response back to []float32.
func ToVector32(s string) []float32 { _ = "STUB: not implemented"; return nil }

// VectorString64 convert the provided []float64 into a string. Users can use this to build vector search queries:
//
//	client.B().FtSearch().Index("idx").Query("*=>[KNN 5 @vec $V]").
//	    Params().Nargs(2).NameValue().NameValue("V", valkey.VectorString64([]float64{1})).
//	    Dialect(2).Build()
func VectorString64(v []float64) string { _ = "STUB: not implemented"; return "" }

// ToVector64 reverts VectorString64. User can use this to convert valkey response back to []float64.
func ToVector64(s string) []float64 { _ = "STUB: not implemented"; return nil }

// JSON convert the provided parameter into a JSON string. Users can use this JSON helper to work with RedisJSON commands.
// For example:
//
//	client.B().JsonSet().Key("a").Path("$.myField").Value(valkey.JSON("str")).Build()
func JSON(in any) string { _ = "STUB: not implemented"; return "" }
