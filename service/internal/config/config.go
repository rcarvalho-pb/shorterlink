package config

import "os"

func GetEnv(prop, fallback string) string {
	property := os.Getenv(prop)
	if property == "" {
		return fallback
	}
	return property
}
