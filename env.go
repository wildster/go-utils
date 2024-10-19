package goutils

import (
	"os"
	"strconv"
)

func DefaultString(s string) *string {
	val := s

	return &val
}

func DeafultInt(s int) *int {
	val := s

	return &val
}

func GetStringEnv(key string, defaultValue *string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if defaultValue != nil {
		return *defaultValue
	}

	panic("%s is required but not set")
}

func GetIntEnv(key string, defaultValue *int) int {
	if value, ok := os.LookupEnv(key); ok {
		n, err := strconv.Atoi(value)

		if err == nil {
			return n
		}
	}

	if defaultValue != nil {
		return *defaultValue
	}

	panic("%s is required but not set")
}

func GetBoolEnv(key string) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)

		if err == nil {
			return b
		}
	}

	return false
}
