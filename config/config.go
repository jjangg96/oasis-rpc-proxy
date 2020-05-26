package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
)

const (
	modeDevelopment = "development"
	modeProduction  = "production"
)

var (
	errOasisSocketRequired = errors.New("oasis socket path is required")
)

// Config holds the configuration data
type Config struct {
	AppEnv             string `json:"app_env" envconfig:"APP_ENV" default:"development"`
	OasisSocket        string `json:"oasis_socket" envconfig:"OASIS_SOCKET"`
	ServerAddr         string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort         int64  `json:"server_port" envconfig:"SERVER_PORT" default:"50051"`
	GenesisFilePath    string `json:"genesis_file_path" envconfig:"GENESIS_FILE_PATH" default:"genesis.json"`
	LogLevel           string `json:"log_level" envconfig:"LOG_LEVEL"`
	LogOutput          string `json:"log_output" envconfig:"LOG_OUTPUT"`
	RollbarAccessToken string `json:"rollbar_access_token" envconfig:"ROLLBAR_ACCESS_TOKEN"`
}

// Validate returns an error if config is invalid
func (c *Config) Validate() error {
	if c.OasisSocket == "" {
		return errOasisSocketRequired
	}

	return nil
}

// IsDevelopment returns true if app is in dev mode
func (c *Config) IsDevelopment() bool {
	return c.AppEnv == modeDevelopment
}

// IsProduction returns true if app is in production mode
func (c *Config) IsProduction() bool {
	return c.AppEnv == modeProduction
}

// ListenAddr returns a full listen address and port
func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}

// New returns a new config
func New() *Config {
	return &Config{}
}

// FromFile reads the config from a file
func FromFile(path string, config *Config) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}

// FromEnv reads the config from environment variables
func FromEnv(config *Config) error {
	return envconfig.Process("", config)
}
