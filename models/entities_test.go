package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Models_Entities_User_TestSuite struct {
	suite.Suite
	testable *User
}

func (suite *Models_Entities_User_TestSuite) SetupTest() {
	suite.testable = &User{}
}

func (suite *Models_Entities_User_TestSuite) TestTableName() {
	assert.Equal(suite.T(), TableUsers, suite.testable.TableName())
}

func (suite *Models_Entities_User_TestSuite) TestBeforeCreate() {
	assert.Empty(suite.T(), suite.testable.CreatedAt)
	err := suite.testable.BeforeCreate(nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), suite.testable.CreatedAt)
}

func (suite *Models_Entities_User_TestSuite) TestBeforeUpdate() {
	assert.Empty(suite.T(), suite.testable.ModifiedAt)
	err := suite.testable.BeforeUpdate(nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), suite.testable.ModifiedAt)
}

func Test_Models_Entities_User_TestSuite(t *testing.T) {
	suite.Run(t, new(Models_Entities_User_TestSuite))
}
