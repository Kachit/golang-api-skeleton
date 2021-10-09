package models

import sq "github.com/Masterminds/squirrel"

type QueryBuilder struct {
}

func (q *QueryBuilder) BuildSelectQuery(table string, columns ...string) sq.SelectBuilder {
	return sq.Select(columns...).
		From(table).
		PlaceholderFormat(sq.Dollar)
}

func (q *QueryBuilder) BuildUpdateQuery(table string, rowMap map[string]interface{}) sq.UpdateBuilder {
	return sq.Update(table).
		SetMap(rowMap).
		PlaceholderFormat(sq.Dollar)
}

func (q *QueryBuilder) BuildInsertQuery(table string, rowMap map[string]interface{}) sq.InsertBuilder {
	return sq.Insert(table).
		SetMap(rowMap).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)
}
