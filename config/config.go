package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
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
	AppEnv           string `json:"app_env" envconfig:"APP_ENV" default:"development"`
	OasisSocket      string `json:"oasis_socket" envconfig:"OASIS_SOCKET"`
	ServerAddr       string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort       int64  `json:"server_port" envconfig:"SERVER_PORT" default:"50051"`
	GenesisFilePath  string `json:"genesis_file_path" envconfig:"GENESIS_FILE_PATH" default:"genesis.json"`
	LogLevel         string `json:"log_level" envconfig:"LOG_LEVEL"`
	LogOutput        string `json:"log_output" envconfig:"LOG_OUTPUT"`
	ServerMetricAddr string `json:"server_metric_addr" envconfig:"SERVER_METRIC_ADDR" default:":8081"`
	MetricServerUrl  string `json:"metric_server_url" envconfig:"METRIC_SERVER_URL" default:"/metrics"`
	GrpcMaxRecvSize  int    `json:"grpc_max_recv_size" envconfig:"GRPC_MAX_RECV_SIZE" default:"1073741824"` // 1024^3
	GrpcMaxSendSize  int    `json:"grpc_max_send_size" envconfig:"GRPC_MAX_SEND_SIZE" default:"1073741824"` // 1024^3

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
