package config

import (
	"errors"
	"os"
	"strings"
)

func Secret(name string) (string, error) {
	v := os.Getenv(name)
	if strings.TrimSpace(v) == "" {
		return "", errors.New("secret not configured: " + name)
	}
	return v, nil
}
func Redact(value string) string {
	if value == "" {
		return ""
	}
	if len(value) <= 4 {
		return "****"
	}
	return value[:2] + "***" + value[len(value)-2:]
}
