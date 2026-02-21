package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mdipanjan/hive-v0/internal/styles"
)

type Config struct {
	Theme string `json:"theme"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "hive", "config.json")
}

func Load() Config {
	cfg := Config{Theme: styles.TokyoNight.Key}
	data, err := os.ReadFile(configPath())
	if err == nil {
		json.Unmarshal(data, &cfg)
	}
	return cfg
}

func Save(cfg Config) error {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0755)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, data, 0644)
}
