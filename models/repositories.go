package models

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	TABLE_USERS string = "users"
)

func NewRepositoriesFactory(database *sqlx.DB) *RepositoriesFactory {
	return &RepositoriesFactory{db: database}
}

type RepositoriesFactory struct {
	db *sqlx.DB
}

func (f *RepositoriesFactory) GetUsersRepository() *UsersRepository {
	return &UsersRepository{RepositoryAbstract: f.GetRepositoryAbstract(TABLE_USERS), Mapper: &UserMapper{}}
}

func (f *RepositoriesFactory) GetRepositoryAbstract(table string) *RepositoryAbstract {
	return &RepositoryAbstract{database: f.db, queryBuilder: &QueryBuilder{}, table: table}
}

type RepositoryAbstract struct {
	database     *sqlx.DB
	queryBuilder *QueryBuilder
	table        string
}

func (ra *RepositoryAbstract) fetchAll(collection interface{}, condition interface{}, limit uint64, offset uint64, orderBy map[string]string) error {
	qb := ra.queryBuilder.BuildSelectQuery(ra.table, "*").Where(condition)

	if limit > 0 {
		qb = qb.Limit(limit)
	}
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	if orderBy != nil {
		for k, v := range orderBy {
			qb = qb.OrderBy(k + " " + v)
		}
	}

	q, v, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("RepositoryAbstract.fetchAll: %v", err)
	}
	return ra.database.Select(collection, q, v...)
}

func (ra *RepositoryAbstract) fetchOne(entity interface{}, condition interface{}) error {
	q, v, _ := ra.queryBuilder.
		BuildSelectQuery(ra.table, "*").
		Where(condition).
		ToSql()

	return ra.database.Get(entity, q, v...)
}

func (ra *RepositoryAbstract) fetchColumn(column string, condition interface{}) *sql.Row {
	q, v, _ := ra.queryBuilder.
		BuildSelectQuery(ra.table, column).
		Where(condition).
		ToSql()

	return ra.database.QueryRow(q, v...)
}

func (ra *RepositoryAbstract) insert(rowMap map[string]interface{}) (uint64, error) {
	query, row, _ := ra.queryBuilder.
		BuildInsertQuery(ra.table, rowMap).
		ToSql()

	var lastInsertId uint64
	err := ra.database.QueryRow(query, row...).Scan(&lastInsertId)
	if err != nil {
		return lastInsertId, fmt.Errorf("RepositoryAbstract.insert: %v", err)
	}
	return lastInsertId, nil
}

func (ra *RepositoryAbstract) update(rowMap map[string]interface{}, pred interface{}) (int64, error) {
	query, row, _ := ra.queryBuilder.BuildUpdateQuery(ra.table, rowMap).
		Where(pred).
		ToSql()

	result, err := ra.database.Exec(query, row...)
	if err != nil {
		return 0, fmt.Errorf("RepositoryAbstract.update: %v", err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return affectedRows, fmt.Errorf("RepositoryAbstract.update: %v", err)
	}
	if affectedRows == 0 {
		return affectedRows, fmt.Errorf("RepositoryAbstract.update: row is not updated")
	}
	return affectedRows, nil
}
