package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`

	// Connection Pool
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`
	MaxLifetime  int `yaml:"max_lifetime"`  // in seconds
	MaxIdleTime  int `yaml:"max_idle_time"` // in seconds
}

var cfg *Config = &Config{}

func LoadConfig(filename string) error {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileBytes, cfg)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}
