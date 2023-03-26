package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestNewConfig() {
	cfg, _ := NewConfig("../config.yml")
	assert.IsType(suite.T(), (*Config)(nil), cfg)
}

func (suite *ConfigTestSuite) TestNewConfigWrongFile() {
	cfg, err := NewConfig("../config-foo.yml")
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), cfg)
}

func (suite *ConfigTestSuite) TestConfigGetAPIPort() {
	cfg, _ := NewConfig("../config.yml")
	assert.Equal(suite.T(), ":8080", cfg.GetAppPort())
}

func (suite *ConfigTestSuite) TestConfigGetDatabaseDsn() {
	cfg, _ := NewConfig("../config.yml")
	cfg.Database.User = "example_user"
	cfg.Database.Name = "example_db"
	cfg.Database.Port = 54321
	assert.Equal(suite.T(), "host=127.0.0.1 port=54321 user=example_user dbname=example_db password=12345 sslmode=disable", cfg.GetDatabaseDsn())
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
