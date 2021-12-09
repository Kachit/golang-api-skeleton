package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Port  uint
	Debug bool
}

type DatabaseConfig struct {
	Host           string
	Port           uint16
	Name           string
	User           string
	Password       string
	MaxConnections int `mapstructure:"max_connections"`
}

func (c *Config) GetAppPort() string {
	return ":" + strconv.FormatUint(uint64(c.App.Port), 10)
}

func NewConfig(configPath string) (*Config, error) {
	dir, err := filepath.Abs(filepath.Dir(configPath))
	if err != nil {
		return nil, errors.Wrap(err, "Parse config path")
	}

	configName := filepath.Base(configPath)
	configNameWithoutExt := strings.TrimSuffix(configName, filepath.Ext(configName))

	viper.AddConfigPath(dir)
	viper.SetConfigName(configNameWithoutExt)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "Reading config file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "Parsing config file")
	}

	return &cfg, nil
}
