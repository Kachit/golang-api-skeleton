package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_Entities_User_BeforeInsert(t *testing.T) {
	entity := User{}
	assert.Empty(t, entity.CreatedAt)
	entity.BeforeInsert()
	assert.NotEmpty(t, entity.CreatedAt)
}
