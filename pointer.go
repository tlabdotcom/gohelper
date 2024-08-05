package gohelper

import (
	"github.com/google/uuid"
)

// PointerUUID returns a pointer to the given UUID.
// It is useful for creating pointers to UUID values.
func PointerUUID(id uuid.UUID) *uuid.UUID {
	return &id
}

// PointerString returns a pointer to the given string.
// It is useful for creating pointers to string values.
func PointerString(prm string) *string {
	return &prm
}

// PointerBool returns a pointer to the given bool.
// It is useful for creating pointers to bool values.
func PointerBool(prm bool) *bool {
	return &prm
}

// PointerInt returns a pointer to the given int.
// It is useful for creating pointers to int values.
func PointerInt(prm int) *int {
	return &prm
}

// PointerInt8 returns a pointer to the given int8.
// It is useful for creating pointers to int8 values.
func PointerInt8(prm int8) *int8 {
	return &prm
}

// PointerInt16 returns a pointer to the given int16.
// It is useful for creating pointers to int16 values.
func PointerInt16(prm int16) *int16 {
	return &prm
}

// PointerInt32 returns a pointer to the given int32.
// It is useful for creating pointers to int32 values.
func PointerInt32(prm int32) *int32 {
	return &prm
}

// PointerInt64 returns a pointer to the given int64.
// It is useful for creating pointers to int64 values.
func PointerInt64(prm int64) *int64 {
	return &prm
}

// PointerUint returns a pointer to the given uint.
// It is useful for creating pointers to uint values.
func PointerUint(prm uint) *uint {
	return &prm
}

// PointerUint8 returns a pointer to the given uint8.
// It is useful for creating pointers to uint8 values.
func PointerUint8(prm uint8) *uint8 {
	return &prm
}

// PointerUint16 returns a pointer to the given uint16.
// It is useful for creating pointers to uint16 values.
func PointerUint16(prm uint16) *uint16 {
	return &prm
}

// PointerUint32 returns a pointer to the given uint32.
// It is useful for creating pointers to uint32 values.
func PointerUint32(prm uint32) *uint32 {
	return &prm
}

// PointerUint64 returns a pointer to the given uint64.
// It is useful for creating pointers to uint64 values.
func PointerUint64(prm uint64) *uint64 {
	return &prm
}

// PointerFloat32 returns a pointer to the given float32.
// It is useful for creating pointers to float32 values.
func PointerFloat32(prm float32) *float32 {
	return &prm
}

// PointerFloat64 returns a pointer to the given float64.
// It is useful for creating pointers to float64 values.
func PointerFloat64(prm float64) *float64 {
	return &prm
}

// PointerComplex64 returns a pointer to the given complex64.
// It is useful for creating pointers to complex64 values.
func PointerComplex64(prm complex64) *complex64 {
	return &prm
}

// PointerComplex128 returns a pointer to the given complex128.
// It is useful for creating pointers to complex128 values.
func PointerComplex128(prm complex128) *complex128 {
	return &prm
}

// PointerByte returns a pointer to the given byte.
// It is useful for creating pointers to byte values.
func PointerByte(prm byte) *byte {
	return &prm
}

// PointerRune returns a pointer to the given rune.
// It is useful for creating pointers to rune values.
func PointerRune(prm rune) *rune {
	return &prm
}
