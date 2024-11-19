package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	OpenAIKey string `json:"openai_key"`
	OllamaServer string `json:"ollama_server,omitempty"`
	DefaultModel string `json:"default_model,omitempty"`
}

func getConfig() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
		os.Exit(1)
	}
	configFilePath := filepath.Join(homeDir, ".config",".ai-commit-gen-config.json")
	log.Println("Config file path:", configFilePath)
	return configFilePath
}

func LoadConfig() Config {
	configFilePath := getConfig()
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Error opening config file", err)
		os.Exit(1)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config file", err)
		os.Exit(1)
	}
	return config
}
