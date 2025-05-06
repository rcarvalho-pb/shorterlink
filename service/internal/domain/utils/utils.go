package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strings"
)

func GenerateSlug(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	slug := base64.URLEncoding.EncodeToString(b)
	slug = strings.ReplaceAll(slug, "-", "")
	slug = strings.ReplaceAll(slug, "_", "")
	slug = strings.ToLower(slug)
	return slug
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
