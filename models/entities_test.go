package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_Entities_User_TableName(t *testing.T) {
	entity := User{}
	assert.Equal(t, TableUsers, entity.TableName())
}

func Test_Models_Entities_User_BeforeCreate(t *testing.T) {
	entity := User{}
	assert.Empty(t, entity.CreatedAt)
	err := entity.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, entity.CreatedAt)
}

func Test_Models_Entities_User_BeforeUpdate(t *testing.T) {
	entity := User{}
	assert.Empty(t, entity.ModifiedAt)
	err := entity.BeforeUpdate(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, entity.ModifiedAt)
}
