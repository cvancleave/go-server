package config

import (
	"encoding/json"
)

const (
	secretName   = "go-server/config"
	secretRegion = "us-east-1"
)

// Config - holds configuration information for the server and integrations
type Config struct {
	// server information
	ServerPort int    `json:"server_port"`
	LogLevel   string `json:"log_level"`

	// auth information
	TokenKey     string `json:"token_key"`
	TokenIssuer  string `json:"token_issuer"`
	TokenTimeout int    `json:"token_timeout"`
}

// Load - load config from AWS secrets manager
func Load() (*Config, error) {

	// secretJson, err := utils.GetSecret(secretName, secretRegion)
	// if err != nil {
	// 	return nil, err
	// }

	// return dummy secret
	secretJson := []byte(`{
		"server_port": 4000,
		"log_level": "debug",
		"token_key": "tokenEncryptionKey",
		"token_issuer": "tokenIssuer",
		"token_timeout": 60
	}`)

	cfg := &Config{}
	if err := json.Unmarshal([]byte(secretJson), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
