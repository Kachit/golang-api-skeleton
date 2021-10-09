package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_Mappers_UserMapper_MapForInsertByDefault(t *testing.T) {
	mapper := UserMapper{}
	user := User{}

	row := map[string]interface{}{
		"name":        user.Name,
		"email":       user.Email,
		"password":    user.Password,
		"created_at":  user.CreatedAt,
		"description": user.Description.String,
	}
	result := mapper.MapForInsert(&user)
	assert.Equal(t, row, result)
}

func Test_Models_Mappers_UserMapper_MapForUpdateByDefault(t *testing.T) {
	mapper := UserMapper{}
	user := User{}

	row := map[string]interface{}{
		"name":        user.Name,
		"email":       user.Email,
		"password":    user.Password,
		"created_at":  user.CreatedAt,
		"description": user.Description.String,
	}
	result := mapper.MapForUpdate(&user)
	assert.Equal(t, row, result)
}
