package infrastructure

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Auth     AuthConfig     `mapstructure:"auth"`
	Crypt    CryptoConfig   `mapstructure:"crypt"`
	HashIds  HashIdsConfig  `mapstructure:"hashids"`
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

type AppConfig struct {
	Port  uint
	Debug bool
}

type HashIdsConfig struct {
	Length int
	Salt   string
}

type CryptoConfig struct {
	Cost int
}

type AuthConfig struct {
	Header  string
	Token   string
	Enabled bool
}

type DatabaseConfig struct {
	Host               string
	Port               uint16
	Name               string
	User               string
	Password           string
	MaxConnections     int    `mapstructure:"max_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
	SslMode            string `mapstructure:"sslmode"`
}

type LoggerConfig struct {
	Mattermost LoggerAdapterMattermostConfig `mapstructure:"mattermost"`
}

type LoggerAdapterMattermostConfig struct {
	WebhookUrl string `mapstructure:"webhook_url"`
	Username   string `mapstructure:"user_name"`
}

func (c *Config) GetAppPort() string {
	return ":" + strconv.FormatUint(uint64(c.App.Port), 10)
}

func (c *Config) GetDatabaseDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Name,
		c.Database.Password,
		c.Database.SslMode,
	)
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

func NewConfigMock() (*Config, error) {
	return NewConfig("../config.yml")
}
