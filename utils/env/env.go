package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func InitEnv(basePath string) error {
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
	for _, allowed := range allowedEnvFiles {
		if basePath == allowed {
			return basePath, nil
		}
	}

	return "", fmt.Errorf("invalid environment file: %s", basePath)
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}
	return value
}

func GetEnvInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic("environment variable not set")
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("error converting environment variable to int")
	}
	return intValue
}
