package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	cfg, _ := NewConfig("../config.yml")
	assert.IsType(t, (*Config)(nil), cfg)
}

func Test_Config_NewConfigWrongFile(t *testing.T) {
	cfg, err := NewConfig("../config-foo.yml")
	assert.Error(t, err)
	assert.Empty(t, cfg)
}

func Test_Config_GetAPIPort(t *testing.T) {
	cfg, _ := NewConfig("../config.yml")
	assert.Equal(t, ":8080", cfg.GetAppPort())
}

func Test_Config_GetDatabaseDsn(t *testing.T) {
	cfg, _ := NewConfig("../config.yml")
	cfg.Database.User = "example_user"
	cfg.Database.Name = "example_db"
	cfg.Database.Port = 54321
	assert.Equal(t, "host=127.0.0.1 port=54321 user=example_user dbname=example_db password=12345 sslmode=disable", cfg.GetDatabaseDsn())
}
