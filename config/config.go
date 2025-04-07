package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nakachan-ing/reflect-cli/model"
	"gopkg.in/yaml.v3"
)

func GetConfigPath() (string, error) {
	// Check if the environment variable `ZK_CONFIG` is set
	if customConfig := os.Getenv("ZK_CONFIG"); customConfig != "" {
		return customConfig, nil
	}

	var configPath string

	switch runtime.GOOS {
	case "windows":
		// Use `APPDATA\reflect-cli\config.yaml` if available
		appData := os.Getenv("APPDATA")
		if appData != "" {
			configPath = filepath.Join(appData, "reflect-cli", "zk-config.yaml")
		} else {
			// Fallback to `USERPROFILE` if `APPDATA` is unavailable
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return "", fmt.Errorf("failed to determine home directory: %w", err)
			}
			configPath = filepath.Join(homeDir, "AppData", "Roaming", "reflect-cli", "zk-config.yaml")
		}

	default: // macOS / Linux
		homeDir, homeErr := os.UserHomeDir()
		if homeErr != nil {
			return "", fmt.Errorf("failed to determine home directory: %w", homeErr)
		}
		configPath = filepath.Join(homeDir, ".config", "reflect-cli", "zk-config.yaml")
		log.Printf("Failed to get user config directory, using fallback: %s", configPath)
	}

	return configPath, nil
}

// Expand `~` to the home directory (Windows included)
func expandHomeDir(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Printf("Failed to get home directory: %v", err)
			return path
		}
		return filepath.Join(home, path[2:])
	}
	return path
}

func LoadConfig() (*model.Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file (%s): %w", configPath, err)
	}

	var config model.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Expand `~` in paths
	config.BaseDir = expandHomeDir(config.BaseDir)
	config.TemplateDir = expandHomeDir(config.TemplateDir)

	return &config, nil
}
