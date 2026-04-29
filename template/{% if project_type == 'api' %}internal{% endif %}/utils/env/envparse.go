package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// String returns the value of the environment variable named by key,
// or fallback if the variable is unset or empty.
func String(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Int returns the environment variable as an int, or fallback if unset/empty.
// Returns an error if the value cannot be parsed.
func Int(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected int, got %q", key, v)
	}
	return n, nil
}

// Bool returns the environment variable as a bool, or fallback if unset/empty.
// Accepted true values: "1", "true", "yes", "on" (case-insensitive).
// Returns an error if the value cannot be parsed.
func Bool(key string, fallback bool) (bool, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	switch strings.ToLower(v) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, fmt.Errorf("env %s: expected bool, got %q", key, v)
	}
}

// Duration returns the environment variable parsed as a time.Duration,
// or fallback if unset/empty. Uses time.ParseDuration format (e.g. "30s", "1m").
// Returns an error if the value cannot be parsed.
func Duration(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected duration, got %q", key, v)
	}
	return d, nil
}

// MustString returns the value of the environment variable named by key.
// Panics if the variable is unset or empty.
func MustString(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %q is not set", key))
	}
	return v
}
