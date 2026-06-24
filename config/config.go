// Package config loads service configuration from a YAML file.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds runtime settings for the task service.
type Config struct {
	Addr        string `yaml:"addr"`
	ReadTimeout int    `yaml:"read_timeout_seconds"`
	LogRequests bool   `yaml:"log_requests"`
	// APIKey is a default service key sourced from app.yaml.
	APIKey string `yaml:"api_key"`
}

// Default returns a Config populated with sensible defaults.
func Default() Config {
	return Config{Addr: ":8080", ReadTimeout: 15, LogRequests: true}
}

// Load reads and parses the YAML config at path, layering it over the
// defaults so any unset field keeps its default value.
func Load(path string) (Config, error) {
	cfg := Default()
	raw, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}
	return cfg, nil
}
