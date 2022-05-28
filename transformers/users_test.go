package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Transformers_UsersTransformer_TestSuite struct {
	suite.Suite
	hashIds  *infrastructure.HashIds
	testable *UsersTransformer
}

func (suite *Transformers_UsersTransformer_TestSuite) SetupTest() {
	cfg, _ := testable.NewConfigMock()
	suite.hashIds = infrastructure.NewHashIds(cfg)
	suite.testable = NewUsersTransformer(suite.hashIds)
}

func (suite *Transformers_UsersTransformer_TestSuite) TestToUserFromStruct() {
	user := models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result := suite.testable.toUser(fractal.Any(user))
	assert.Equal(suite.T(), user.Id, result.Id)
	assert.Equal(suite.T(), user.Name, result.Name)
	assert.Equal(suite.T(), user.Email, result.Email)
	assert.Equal(suite.T(), user.Password, result.Password)
}

func (suite *Transformers_UsersTransformer_TestSuite) TestToUserFromPointer() {
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result := suite.testable.toUser(fractal.Any(user))
	assert.Equal(suite.T(), user.Id, result.Id)
	assert.Equal(suite.T(), user.Name, result.Name)
	assert.Equal(suite.T(), user.Email, result.Email)
	assert.Equal(suite.T(), user.Password, result.Password)
}

func (suite *Transformers_UsersTransformer_TestSuite) TestTransformUserFull() {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 10:10:10")
	modifiedAt, _ := time.Parse("2006-01-02 15:04:05", "2021-02-01 10:10:10")
	deletedAt, _ := time.Parse("2006-01-02 15:04:05", "2021-03-01 10:10:10")
	da := gorm.DeletedAt{Time: deletedAt}
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd", CreatedAt: createdAt, ModifiedAt: &modifiedAt, DeletedAt: da}
	result := suite.testable.Transform(user)
	assert.Equal(suite.T(), "ngB0NV05ev", result["id"])
	assert.Equal(suite.T(), user.Name, result["name"])
	assert.Equal(suite.T(), user.Email, result["email"])
	assert.Equal(suite.T(), createdAt, result["created_at"])
	assert.Equal(suite.T(), &modifiedAt, result["modified_at"])
	assert.Equal(suite.T(), deletedAt, result["deleted_at"])
}

func (suite *Transformers_UsersTransformer_TestSuite) TestTransformUserSimple() {
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result := suite.testable.Transform(user)
	assert.Equal(suite.T(), "ngB0NV05ev", result["id"])
	assert.Equal(suite.T(), user.Name, result["name"])
	assert.Equal(suite.T(), user.Email, result["email"])
}

func (suite *Transformers_UsersTransformer_TestSuite) TestTransformAnotherStruct() {
	user := &testable.StubUser{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result := suite.testable.Transform(user)
	assert.Empty(suite.T(), result)
}

func (suite *Transformers_UsersTransformer_TestSuite) TestTransformUsersToFractal() {
	users := []*models.User{{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}}
	result := transformUsersToFractal(users)
	fractalUser := result[0].(models.User)
	assert.Equal(suite.T(), users[0].Id, fractalUser.Id)
	assert.Equal(suite.T(), users[0].Name, fractalUser.Name)
	assert.Equal(suite.T(), users[0].Email, fractalUser.Email)
	assert.Equal(suite.T(), users[0].Password, fractalUser.Password)
}

func Test_Transformers_UsersTransformer_TestSuite(t *testing.T) {
	suite.Run(t, new(Transformers_UsersTransformer_TestSuite))
}
