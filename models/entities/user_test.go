package entities

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ModelsEntitiesUserTestSuite struct {
	suite.Suite
	testable *User
}

func (suite *ModelsEntitiesUserTestSuite) SetupTest() {
	suite.testable = &User{}
}

func (suite *ModelsEntitiesUserTestSuite) TestTableName() {
	assert.Equal(suite.T(), TableUsers, suite.testable.TableName())
}

func (suite *ModelsEntitiesUserTestSuite) TestBeforeCreate() {
	assert.Empty(suite.T(), suite.testable.CreatedAt)
	err := suite.testable.BeforeCreate(nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), suite.testable.CreatedAt)
}

func (suite *ModelsEntitiesUserTestSuite) TestBeforeUpdate() {
	assert.Empty(suite.T(), suite.testable.ModifiedAt)
	err := suite.testable.BeforeUpdate(nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), suite.testable.ModifiedAt)
}

func TestModelsEntitiesUserTestSuite(t *testing.T) {
	suite.Run(t, new(ModelsEntitiesUserTestSuite))
}
