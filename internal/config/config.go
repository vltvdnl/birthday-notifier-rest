package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Client struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retries"`
}
type ClientsConfig struct {
	SSO Client `yaml:"sso"`
}

type Config struct {
	Env           string `yaml:"env" env-default:"local"`
	StoragePath   string `yaml:"storage_path" env-required:"true"`
	HTTPServer    `yaml:"http_server"`
	ClientsConfig `yaml:"clients"`
}

func MustLoad() *Config {
	cfgPath := fetchCfgPath()
	if cfgPath == "" {
		panic("no config")
	}

	return MustLoadPath(cfgPath)
}
func MustLoadPath(cfgPath string) *Config {
	if _, err := os.Stat(cfgPath); err != nil {
		panic("no config with this path" + err.Error())
	}
	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("cannot read cfg file" + err.Error())
	}
	return &cfg
}
func fetchCfgPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
