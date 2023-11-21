package main

import (
	"fmt"
	"os"
)

func SetEnv(key, value string) (string, error) {
	// Set the environment variable.
	if err := os.Setenv(key, value); err != nil {
		return value, fmt.Errorf("fail to set environment variable %v", err)
	}
	return value, nil
}

func UnsetEnv(key string) error {
	if err := os.Unsetenv(key); err != nil {
		return fmt.Errorf("failed to unset the env key %v", err)
	}
	return nil
}

func GetenvDefault(key, defValue string) (string, error) {
	conStr, ok := os.LookupEnv(key)
	if !ok {
		return defValue, fmt.Errorf("%s not found", key)
	}
	return conStr, nil
}
