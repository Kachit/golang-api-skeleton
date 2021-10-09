package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Models_QueryBuilder_BuildSelectQuery(t *testing.T) {
	builder := QueryBuilder{}
	q, _, _ := builder.BuildSelectQuery("foo", "*").ToSql()
	assert.Equal(t, "SELECT * FROM foo", q)
}

func Test_Models_QueryBuilder_BuildInsertQuery(t *testing.T) {
	builder := QueryBuilder{}
	row := map[string]interface{}{
		"id":   1,
		"code": "foo",
	}
	q, _, _ := builder.BuildInsertQuery("foo", row).ToSql()
	assert.Equal(t, "INSERT INTO foo (code,id) VALUES ($1,$2) RETURNING id", q)
}

func Test_Models_QueryBuilder_BuildUpdateQuery(t *testing.T) {
	builder := QueryBuilder{}
	row := map[string]interface{}{
		"id":   1,
		"code": "foo",
	}
	q, _, _ := builder.BuildUpdateQuery("foo", row).ToSql()
	assert.Equal(t, "UPDATE foo SET code = $1, id = $2", q)
}
