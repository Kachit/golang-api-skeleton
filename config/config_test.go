package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	cfg, _ := NewConfig("../config.yml")
	assert.IsType(t, (*Config)(nil), cfg)
}

func Test_Config_GetAPIPort(t *testing.T) {
	cfg, _ := NewConfig("../config.yml")
	assert.Equal(t, ":8080", cfg.GetAppPort())
}
