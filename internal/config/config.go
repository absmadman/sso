package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

const (
	envConfigPathVariableName = "CONFIG_PATH"
	defaultUsage              = "path to config file"
)

type Config struct {
	Env         string        `yaml:"env" env-required:"true"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPC          `yaml:"grpc"`
}

type GRPC struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

// MustRead read a config file from env "CONFIG_PATH" or from flag "config_path"
func MustRead() *Config {
	path := getConfigPath()
	if path == "" {
		panic("path to config file is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exist")
	}

	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic("some variables in config file is not set")
	}
	return cfg
}

// getConfigPath get a config path first try to parse flag "config_path" then from env variable
// if config_path is not set it takes default value
func getConfigPath() string {
	var res string
	// -config_path=res
	flag.StringVar(&res, "config_path", "", defaultUsage)
	flag.Parse()
	if res == "" {
		// CONFIG_PATH=res
		res = os.Getenv(envConfigPathVariableName)
	}
	return res
}
