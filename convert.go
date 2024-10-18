package gohelper

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// StringToInt converts a string to an int
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// IntToString converts an int to a string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToFloat64 converts a string to a float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float64ToString converts a float64 to a string
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToBool converts a string to a bool
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// BoolToString converts a bool to a string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// StringToTime converts a string to a time.Time
// The layout parameter specifies the format of the input string
func StringToTime(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

// TimeToString converts a time.Time to a string
// The layout parameter specifies the desired output format
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// IntToFloat64 converts an int to a float64
func IntToFloat64(i int) float64 {
	return float64(i)
}

// Float64ToInt converts a float64 to an int
func Float64ToInt(f float64) int {
	return int(f)
}

// StringToInt64 converts a string to an int64
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// Int64ToString converts an int64 to a string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// IntToInt64 converts an int to an int64
func IntToInt64(i int) int64 {
	return int64(i)
}

// Int64ToInt converts an int64 to an int
func Int64ToInt(i int64) int {
	return int(i)
}

// StringToUUID converts a string to a UUID
func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// UUIDToString converts a UUID to a string
func UUIDToString(u uuid.UUID) string {
	return u.String()
}

// NewUUID generates a new UUID
func NewUUID() uuid.UUID {
	return uuid.New()
}

// Interface conversions
func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func ToInt(v interface{}) (int, error) {
	switch value := v.(type) {
	case int:
		return value, nil
	case float64:
		return int(value), nil
	case string:
		return StringToInt(value)
	default:
		return 0, fmt.Errorf("unsupported type for conversion to int")
	}
}

func ToFloat64(v interface{}) (float64, error) {
	switch value := v.(type) {
	case float64:
		return value, nil
	case int:
		return float64(value), nil
	case string:
		return StringToFloat64(value)
	default:
		return 0, fmt.Errorf("unsupported type for conversion to float64")
	}
}

func ToBool(v interface{}) (bool, error) {
	switch value := v.(type) {
	case bool:
		return value, nil
	case string:
		return StringToBool(value)
	case int:
		return value != 0, nil
	default:
		return false, fmt.Errorf("unsupported type for conversion to bool")
	}
}

func ToInt64(v interface{}) (int64, error) {
	switch value := v.(type) {
	case int64:
		return value, nil
	case int:
		return int64(value), nil
	case float64:
		return int64(value), nil
	case string:
		return StringToInt64(value)
	default:
		return 0, fmt.Errorf("unsupported type for conversion to int64")
	}
}

func ToUUID(v interface{}) (uuid.UUID, error) {
	switch value := v.(type) {
	case uuid.UUID:
		return value, nil
	case string:
		return StringToUUID(value)
	default:
		return uuid.Nil, fmt.Errorf("unsupported type for conversion to UUID")
	}
}
