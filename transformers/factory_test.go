package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransformersFactoryTestSuite struct {
	suite.Suite
	fractal  *fractal.Manager
	hashIds  *infrastructure.HashIds
	testable *Factory
}

func (suite *TransformersFactoryTestSuite) SetupTest() {
	cfg, _ := infrastructure.NewConfigMock()
	suite.fractal = infrastructure.NewFractalManager()
	suite.hashIds = infrastructure.NewHashIds(cfg)
	suite.testable = NewTransformersFactory(suite.fractal, suite.hashIds)
}

func (suite *TransformersFactoryTestSuite) TestMapUsersResourceItem() {
	user := &entities.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result, err := suite.testable.MapUsersResourceItem(user)
	resultMap := result.(map[string]interface{})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "ngB0NV05ev", resultMap["id"])
	assert.Equal(suite.T(), user.Name, resultMap["name"])
	assert.Equal(suite.T(), user.Email, resultMap["email"])
}

func (suite *TransformersFactoryTestSuite) TestMapUsersResourceCollection() {
	user := &entities.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result, err := suite.testable.MapUsersResourceCollection([]*entities.User{user})
	resultMap := result.([]interface{})
	resultMapEl := resultMap[0].(map[string]interface{})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "ngB0NV05ev", resultMapEl["id"])
	assert.Equal(suite.T(), user.Name, resultMapEl["name"])
	assert.Equal(suite.T(), user.Email, resultMapEl["email"])
}

func TestTransformersFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransformersFactoryTestSuite))
}
