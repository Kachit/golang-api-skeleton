package transformers

import (
	"fmt"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Transformers_UsersTransformer_MapUsersResourceItem(t *testing.T) {
	factory := NewTransformersFactory(infrastructure.NewFractalManager(), infrastructure.NewHashIds())
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result, err := factory.MapUsersResourceItem(user)
	resultMap := result.(map[string]interface{})
	assert.NoError(t, err)
	assert.Equal(t, "ngB0NV05ev", resultMap["id"])
	assert.Equal(t, user.Name, resultMap["name"])
	assert.Equal(t, user.Email, resultMap["email"])
}

func Test_Transformers_UsersTransformer_MapUsersResourceCollection(t *testing.T) {
	factory := NewTransformersFactory(infrastructure.NewFractalManager(), infrastructure.NewHashIds())
	user := &models.User{Id: 1, Name: "name", Email: "foo@bar.baz", Password: "pwd"}
	result, err := factory.MapUsersResourceCollection([]*models.User{user})
	fmt.Println(result)
	resultMap := result.([]interface{})
	resultMapEl := resultMap[0].(map[string]interface{})
	assert.NoError(t, err)
	assert.Equal(t, "ngB0NV05ev", resultMapEl["id"])
	assert.Equal(t, user.Name, resultMapEl["name"])
	assert.Equal(t, user.Email, resultMapEl["email"])
}
