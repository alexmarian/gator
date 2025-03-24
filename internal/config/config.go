package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(currentUserName string) error {
	c.CurrentUserName = currentUserName
	return write(*c)
}

func Read() (*Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return &Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}
func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
func getConfigFilePath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return userHomeDir + "/" + configFileName, nil
}
