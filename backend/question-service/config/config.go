package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type (
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		GRPC  `yaml:"grpc"`
		Log   `yaml:"logger"`
		Redis `yaml:"redis"`
		Env   string `yaml:"env" env-required:"true" env-default:"local"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port     string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		CertFile string `env-required:"true" yaml:"cert_file" env:"HTTP_CERT_FILE"`
		KeyFile  string `env-required:"true" yaml:"key_file" env:"HTTP_KEY_FILE"`
	}

	GRPC struct {
		Address           string `env-required:"true" yaml:"addr" env:"GRPC_ADDRESS"`
		ConnectionAddress string `env-required:"true" yaml:"conn" env:"GRPC_CONNECTION"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Redis struct {
		Addr string `env-required:"true" yaml:"addr" env:"REDIS_ADDR"`
	}
)

func GetConfig() *Config {
	pathToConfig := fetchConfigPath()
	if _, err := os.Stat(pathToConfig); os.IsNotExist(err) {
		log.Fatalf("file does not exist: %v", err)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(pathToConfig, &cfg); err != nil {
		log.Fatalf("can't parse config: %v", err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
	flag.Parse()
	return configPath
}
