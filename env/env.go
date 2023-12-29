package env

import (
	"os"
	"strconv"
)

// Get uses the os.Getenv() to get environment variable.
func Get(key string) string {
	return os.Getenv(key)
}

// GetDefault overides os.Getenv() with a default value case the key is empty.
func GetDefault(key, defaultValue string) string {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue
	}
	return envrionmentVariable
}

// GetInt uses the os.Getenv() to get environment variable converting it to integer.
func GetInt(key string, defaultValue int) (int, error) {
	return strconv.Atoi(os.Getenv(key))
}

// GetIntDefault overides os.Getenv() with a default value case the key is empty and return a integer value.
func GetIntDefault(key string, defaultValue int) (int, error) {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(envrionmentVariable)
}

// GetMustInt uses the os.Getenv() to get environment variable converting it to integer and panic if it is not integer.
func GetMustInt(key string, defaultValue int) int {
	convertedVariable, _ := strconv.Atoi(os.Getenv(key))
	return convertedVariable
}

// GetMustIntDefault overides os.Getenv() with a default value case the key is empty and return a integer value and panic if it is not integer.
func GetMustIntDefault(key string, defaultValue int) int {
	envrionmentVariable := os.Getenv(key)
	if envrionmentVariable == "" {
		return defaultValue
	}
	convertedVariable, _ := strconv.Atoi(envrionmentVariable)
	return convertedVariable
}
