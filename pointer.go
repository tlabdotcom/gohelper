package gohelper

import (
	"time"

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

// PointerTime returns a pointer to the given time.
// It is useful for creating pointers to time values.
func PointerTime(prm time.Time) *time.Time {
	return &prm
}

// GetUUIDValue returns the value of a uuid.UUID pointer.
func GetUUIDValue(ptr *uuid.UUID) uuid.UUID {
	if ptr == nil {
		return uuid.UUID{}
	}
	return *ptr
}

// GetStringValue returns the value of a string pointer.
func GetStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// GetBoolValue returns the value of a bool pointer.
func GetBoolValue(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return *ptr
}

// GetIntValue returns the value of an int pointer.
func GetIntValue(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetInt8Value returns the value of an int8 pointer.
func GetInt8Value(ptr *int8) int8 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetInt16Value returns the value of an int16 pointer.
func GetInt16Value(ptr *int16) int16 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetInt32Value returns the value of an int32 pointer.
func GetInt32Value(ptr *int32) int32 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetInt64Value returns the value of an int64 pointer.
func GetInt64Value(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetUintValue returns the value of a uint pointer.
func GetUintValue(ptr *uint) uint {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetUint8Value returns the value of a uint8 pointer.
func GetUint8Value(ptr *uint8) uint8 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetUint16Value returns the value of a uint16 pointer.
func GetUint16Value(ptr *uint16) uint16 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetUint32Value returns the value of a uint32 pointer.
func GetUint32Value(ptr *uint32) uint32 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetUint64Value returns the value of a uint64 pointer.
func GetUint64Value(ptr *uint64) uint64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetFloat32Value returns the value of a float32 pointer.
func GetFloat32Value(ptr *float32) float32 {
	if ptr == nil {
		return 0.0
	}
	return *ptr
}

// GetFloat64Value returns the value of a float64 pointer.
func GetFloat64Value(ptr *float64) float64 {
	if ptr == nil {
		return 0.0
	}
	return *ptr
}

// GetComplex64Value returns the value of a complex64 pointer.
func GetComplex64Value(ptr *complex64) complex64 {
	if ptr == nil {
		return 0 + 0i
	}
	return *ptr
}

// GetComplex128Value returns the value of a complex128 pointer.
func GetComplex128Value(ptr *complex128) complex128 {
	if ptr == nil {
		return 0 + 0i
	}
	return *ptr
}

// GetByteValue returns the value of a byte pointer.
func GetByteValue(ptr *byte) byte {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetRuneValue returns the value of a rune pointer.
func GetRuneValue(ptr *rune) rune {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// GetRuneValue returns the value of a rune pointer.
func GetTimeValue(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}
