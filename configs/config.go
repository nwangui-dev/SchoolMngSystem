package configs

import (
	//"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
}

var AppConfig *Config

func LoadConfig() *Config {
	cfg := &Config{}
	
	// Default fallbacks
	cfg.App.Port = os.Getenv("APP_PORT")
	if cfg.App.Port == "" {
		cfg.App.Port = "8080"
	}

	cfg.DB.Host = getEnv("DB_HOST", "localhost")
	cfg.DB.Port = getEnv("DB_PORT", "5432")
	cfg.DB.User = getEnv("DB_USER", "postgres")
	cfg.DB.Password = getEnv("DB_PASSWORD", "postgres")
	cfg.DB.Name = getEnv("DB_NAME", "school_db")
	cfg.DB.SSLMode = getEnv("DB_SSLMODE", "disable")

	file, err := os.Open("configs/configs.yml")
	if err == nil {
		defer file.Close()
		decoder := yaml.NewDecoder(file)
		_ = decoder.Decode(cfg)
	}

	AppConfig = cfg
	return cfg
}

func GetDefaultConfig() *Config {
	cfg := &Config{}
	cfg.App.Port = "8080"
	cfg.DB.Host = "localhost"
	cfg.DB.Port = "5432"
	cfg.DB.User = "postgres"
	cfg.DB.Password = "postgres"
	cfg.DB.Name = "school_db"
	cfg.DB.SSLMode = "disable"
	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}