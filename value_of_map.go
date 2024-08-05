package gohelper

import (
	"fmt"
	"reflect"

	"github.com/labstack/gommon/log"
)

// SafeGetStringValue safely retrieves a string value from a map by key.
func SafeGetStringValue(m map[string]interface{}, key string, defaultValue string) string {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	switch v := value.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		log.Errorf("Key %s has a value of unexpected type %T, returning default value: %v\n", key, value, defaultValue)
		return defaultValue
	}
}

// SafeGetIntValue safely retrieves an int value from a map by key.
func SafeGetIntValue(m map[string]interface{}, key string, defaultValue int) int {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	switch v := value.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		log.Errorf("Key %s has a value of unexpected type %T, returning default value: %v\n", key, value, defaultValue)
		return defaultValue
	}
}

// SafeGetInt64Value safely retrieves an int64 value from a map by key.
func SafeGetInt64Value(m map[string]interface{}, key string, defaultValue int64) int64 {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	switch v := value.(type) {
	case int64:
		return v
	case int:
		return int64(v)
	case float64:
		return int64(v)
	default:
		log.Errorf("Key %s has a value of unexpected type %T, returning default value: %v\n", key, value, defaultValue)
		return defaultValue
	}
}

// SafeGetFloat64Value safely retrieves a float64 value from a map by key.
func SafeGetFloat64Value(m map[string]interface{}, key string, defaultValue float64) float64 {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	switch v := value.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		log.Errorf("Key %s has a value of unexpected type %T, returning default value: %v\n", key, value, defaultValue)
		return defaultValue
	}
}

// SafeGetBoolValue safely retrieves a bool value from a map by key.
func SafeGetBoolValue(m map[string]interface{}, key string, defaultValue bool) bool {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	switch v := value.(type) {
	case bool:
		return v
	default:
		log.Errorf("Key %s has a value of unexpected type %T, returning default value: %v\n", key, value, defaultValue)
		return defaultValue
	}
}

// SafeGetValue safely retrieves a value of any type from a map by key.
func SafeGetValue(m map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, exists := m[key]
	if !exists {
		log.Errorf("Key %s not found, returning default value: %v\n", key, defaultValue)
		return defaultValue
	}

	if reflect.TypeOf(value) != reflect.TypeOf(defaultValue) {
		log.Errorf("Key %s has a value of unexpected type %T, expected type %T, returning default value: %v\n", key, value, defaultValue, defaultValue)
		return defaultValue
	}

	return value
}
