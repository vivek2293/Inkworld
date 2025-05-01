package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cast"
)

// ReadEnv reads environment variables from a file and sets them in the current process.
func ReadEnv(basePath string) error {
	getSanitizedFilePath, err := sanitizeEnvFilePath(basePath)
	if err != nil {
		return err
	}

	getSanitizedFilePath = filepath.Clean(getSanitizedFilePath)
	file, err := os.Open(getSanitizedFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if os.Getenv(key) == "" {
			err := os.Setenv(key, value)
			if err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func sanitizeEnvFilePath(basePath string) (string, error) {
	// List of allowed filenames
	allowedEnvFiles := []string{
		".env",
		".env.prod",
		".env.dev",
	}
	// Check if the filename is in the allowed list
	if slices.Contains(allowedEnvFiles, basePath) {
		return basePath, nil
	}

	return "", fmt.Errorf("invalid environment file: %s", basePath)
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	intValue := cast.ToInt(os.Getenv(key))
	return intValue
}
