package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
	"github.com/kachit/golang-api-skeleton/testable"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_Transformers_UsersTransformer_ToUserFromStruct(t *testing.T) {
	user := models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	transformer := NewUsersTransformer(nil)
	result := transformer.toUser(fractal.Any(user))
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
}

func Test_Transformers_UsersTransformer_ToUserFromPointer(t *testing.T) {
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	transformer := NewUsersTransformer(nil)
	result := transformer.toUser(fractal.Any(user))
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
}

func Test_Transformers_UsersTransformer_TransformUserFull(t *testing.T) {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 10:10:10")
	modifiedAt, _ := time.Parse("2006-01-02 15:04:05", "2021-02-01 10:10:10")
	deletedAt, _ := time.Parse("2006-01-02 15:04:05", "2021-03-01 10:10:10")
	da := gorm.DeletedAt{Time: deletedAt}
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd", CreatedAt: createdAt, ModifiedAt: &modifiedAt, DeletedAt: da}
	transformer := NewUsersTransformer(infrastructure.NewHashIds())
	result := transformer.Transform(user)
	assert.Equal(t, "ngB0NV05ev", result["id"])
	assert.Equal(t, user.Name, result["name"])
	assert.Equal(t, user.Email, result["email"])
	assert.Equal(t, createdAt, result["created_at"])
	assert.Equal(t, &modifiedAt, result["modified_at"])
	assert.Equal(t, deletedAt, result["deleted_at"])
}

func Test_Transformers_UsersTransformer_TransformUserSimple(t *testing.T) {
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	transformer := NewUsersTransformer(infrastructure.NewHashIds())
	result := transformer.Transform(user)
	assert.Equal(t, "ngB0NV05ev", result["id"])
	assert.Equal(t, user.Name, result["name"])
	assert.Equal(t, user.Email, result["email"])
}

func Test_Transformers_UsersTransformer_TransformAnotherStruct(t *testing.T) {
	user := &testable.StubUser{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	transformer := NewUsersTransformer(infrastructure.NewHashIds())
	result := transformer.Transform(user)
	assert.Empty(t, result)
}

func Test_Transformers_TransformUsersToFractal(t *testing.T) {
	users := []*models.User{{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}}
	result := transformUsersToFractal(users)
	fractalUser := result[0].(models.User)
	assert.Equal(t, users[0].Id, fractalUser.Id)
}
