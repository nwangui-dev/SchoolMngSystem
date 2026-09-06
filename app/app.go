package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nwangui-dev/SchoolMngSystem/configs"
	"github.com/nwangui-dev/SchoolMngSystem/db"
	"github.com/nwangui-dev/SchoolMngSystem/server"

	"gopkg.in/yaml.v2"
)

func Start() error {
	configs.LoadConfig()
	db.InitGormDB()

	srv := server.NewServer(configs.AppConfig)
	srv.Start()

	return nil
}

func WriteDefaultConfig(configPath string) error {
	defaultConfig := configs.GetDefaultConfig()

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal default config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Printf("Default configuration written to %s\n", configPath)
	return nil
}

func InitDatabase() error {
	fmt.Println("Initializing database connection and running migrations...")
	configs.LoadConfig()
	db.InitGormDB()
	fmt.Println("Database initialized and migrated successfully.")
	return nil
}