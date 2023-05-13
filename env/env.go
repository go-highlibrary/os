package env

import (
	"os"
	"strconv"
)

// Get overides os.Getenv() with a default value case the key is empty.
func Get(key, defaultValue string) string {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue
	}
	return envrionmentVariable
}

// GetInt overides os.Getenv() with a default value case the key is empty and return a int value.
func GetInt(key string, defaultValue int) (int, error) {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(envrionmentVariable)
}

// GetShouldInt overides os.Getenv() with a default value case the key is empty and return a int value.
func GetShouldInt(key string, defaultValue int) int {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue
	}
	convertedVariable, _ := strconv.Atoi(envrionmentVariable)
	return convertedVariable
}
