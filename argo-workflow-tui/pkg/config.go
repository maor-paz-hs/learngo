package config

import (
	"errors"
	"os"
)

type Config struct {
	ArgoServerURL string
	ArgoToken     string
	ArgoIDToken   string
}

func LoadConfig() *Config {
	cfg := &Config{
		ArgoServerURL: getEnvWithDefault("ARGO_SERVER_URL", "https://argoworkflows-prod.omcomcom.com"),
		ArgoToken:     os.Getenv("ARGO_TOKEN"),
		ArgoIDToken:   os.Getenv("ARGO_ID_TOKEN"),
	}
	return cfg
}

func (c *Config) Validate() error {
	if c.ArgoServerURL == "" {
		return errors.New("ARGO_SERVER_URL is required")
	}

	if c.ArgoToken == "" && c.ArgoIDToken == "" {
		return errors.New("ARGO_TOKEN or ARGO_ID_TOKEN is required")
	}

	return nil
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
