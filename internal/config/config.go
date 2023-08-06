package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Mastodon struct {
		Server       string `json:"Server`
		ClientID     string `json:ClientID`
		ClientSecret string `json:ClientSecret`
		Email        string `json:Email`
		Password     string `json:Password`
	} `json:"mastodon"`
}

var config Config

func Parse() Config {
	envVarName := "CONFIG"
	filename := "config.dev.json"

	data := []byte(os.Getenv(envVarName))

	if len(data) == 0 {
		fmt.Printf("Environment variable %s does not exist. Try to read file %s instead.\n",
			envVarName,
			filename)
		fileData, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Error reading file %s: %v\n", filename, err)
		}
		data = fileData
	}

	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	if config.Mastodon.Server == "" ||
		config.Mastodon.ClientID == "" ||
		config.Mastodon.ClientSecret == "" {
		log.Fatal("Failed to parse config: fields are missing.")
	}

	return config
}
