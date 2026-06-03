package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadStructFromJsonFile(file string, dst any) error {
	f, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to load %s, err: %w", file, err)
	}

	if err := json.Unmarshal(f, &dst); err != nil {
		return fmt.Errorf("failed to parse %s file, err: %w", file, err)
	}

	return nil
}

func ReadEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}
