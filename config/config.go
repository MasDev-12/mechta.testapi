package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbSetting      DbSetting      `json:"database"`
	ServerSetting  ServerSetting  `json:"rest_server"`
	Argon2Setting  Argon2Setting  `json:"argon2"`
	SwaggerSetting SwaggerSetting `json:"swagger"`
}

type ServerSetting struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DbSetting struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	User                  string `json:"user"`
	Password              string `json:"password"`
	DBName                string `json:"dbname"`
	SSLMode               string `json:"sslmode"`
	MaxConnections        int    `json:"max_connections"`
	MinConnections        int    `json:"min_connections"`
	MaxIdleConnections    int    `json:"max_idle_connections"`
	ConnectionMaxLifetime int    `json:"connection_max_lifetime"`  // in seconds
	ConnectionMaxIdleTime int    `json:"connection_max_idle_time"` // in seconds
	ConnectTimeout        int    `json:"connect_timeout"`          // in seconds
	ReadTimeout           int    `json:"read_timeout"`             // in seconds
	WriteTimeout          int    `json:"write_timeout"`            // in seconds
	ApplicationName       string `json:"application_name"`
}

type Argon2Setting struct {
	Memory     uint32 `json:"memory"`
	Time       uint32 `json:"time"`
	Threads    uint8  `json:"threads"`
	KeyLength  uint32 `json:"keyLength"`
	SaltLength uint32 `json:"saltLength"`
}

type SwaggerSetting struct {
	Host         string `json:"host"`
	Description  string `json:"description"`
	PageTitle    string `json:"pageTitle"`
	Version      string `json:"version"`
	BasePath     string `json:"base_path"`
	ContactName  string `json:"contact_name"`
	ContactURL   string `json:"contact_url"`
	ContactEmail string `json:"contact_email"`
}

func LoadSettingsDb(filePath string) (*DbSetting, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config.DbSetting, nil
}

func LoadSettingArgon2(filePath string) (*Argon2Setting, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config.Argon2Setting, nil
}

func LoadSettingServer(filePath string) (*ServerSetting, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config.ServerSetting, nil
}
func LoadSettingsSwagger(filePath string) (*SwaggerSetting, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &config.SwaggerSetting, nil
}
