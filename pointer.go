package gohelper

import (
	"github.com/google/uuid"
)

// PointerUUID returns a pointer to the given UUID.
func PointerUUID(id uuid.UUID) *uuid.UUID {
	return &id
}

// PointerString returns a pointer to the given string.
func PointerString(prm string) *string {
	return &prm
}

// PointerBool returns a pointer to the given bool.
func PointerBool(prm bool) *bool {
	return &prm
}

// PointerInt returns a pointer to the given int.
func PointerInt(prm int) *int {
	return &prm
}

// PointerInt8 returns a pointer to the given int8.
func PointerInt8(prm int8) *int8 {
	return &prm
}

// PointerInt16 returns a pointer to the given int16.
func PointerInt16(prm int16) *int16 {
	return &prm
}

// PointerInt32 returns a pointer to the given int32.
func PointerInt32(prm int32) *int32 {
	return &prm
}

// PointerInt64 returns a pointer to the given int64.
func PointerInt64(prm int64) *int64 {
	return &prm
}

// PointerUint returns a pointer to the given uint.
func PointerUint(prm uint) *uint {
	return &prm
}

// PointerUint8 returns a pointer to the given uint8.
func PointerUint8(prm uint8) *uint8 {
	return &prm
}

// PointerUint16 returns a pointer to the given uint16.
func PointerUint16(prm uint16) *uint16 {
	return &prm
}

// PointerUint32 returns a pointer to the given uint32.
func PointerUint32(prm uint32) *uint32 {
	return &prm
}

// PointerUint64 returns a pointer to the given uint64.
func PointerUint64(prm uint64) *uint64 {
	return &prm
}

// PointerFloat32 returns a pointer to the given float32.
func PointerFloat32(prm float32) *float32 {
	return &prm
}

// PointerFloat64 returns a pointer to the given float64.
func PointerFloat64(prm float64) *float64 {
	return &prm
}

// PointerComplex64 returns a pointer to the given complex64.
func PointerComplex64(prm complex64) *complex64 {
	return &prm
}

// PointerComplex128 returns a pointer to the given complex128.
func PointerComplex128(prm complex128) *complex128 {
	return &prm
}

// PointerByte returns a pointer to the given byte.
func PointerByte(prm byte) *byte {
	return &prm
}

// PointerRune returns a pointer to the given rune.
func PointerRune(prm rune) *rune {
	return &prm
}
