package models

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type UsersRepository struct {
	*RepositoryAbstract
	Mapper *UserMapper
}

func (pr *UsersRepository) GetList(condition interface{}, limit uint64, offset uint64, orderBy map[string]string) (*UsersCollection, error) {
	collection := &UsersCollection{}
	err := pr.fetchAll(collection, condition, limit, offset, orderBy)
	if err != nil {
		return nil, fmt.Errorf("UsersRepository.GetList: %v", err)
	}
	return collection, nil
}

func (pr *UsersRepository) Count(condition interface{}) (int64, error) {
	row := pr.fetchColumn("COUNT(*)", condition)
	var result int64
	err := row.Scan(&result)
	if err != nil {
		return 0, fmt.Errorf("UsersRepository.Count: %v", err)
	}
	return result, nil
}

func (pr *UsersRepository) GetByID(id uint64) (*User, error) {
	user := &User{}
	err := pr.fetchOne(user, sq.Eq{"id": id})
	if err != nil {
		return nil, fmt.Errorf("UsersRepository.GetByID: %v", err)
	}
	return user, nil
}

func (pr *UsersRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	err := pr.fetchOne(user, sq.Eq{"email": email})
	if err != nil {
		return nil, fmt.Errorf("UsersRepository.GetByEmail: %v", err)
	}
	return user, nil
}

func (pr *UsersRepository) Insert(user *User) (*User, error) {
	rowMap := pr.Mapper.MapForInsert(user)
	lastInsertId, err := pr.insert(rowMap)
	user.Id = lastInsertId
	return user, err
}

func (pr *UsersRepository) Update(user *User) (*User, error) {
	rowMap := pr.Mapper.MapForUpdate(user)
	_, err := pr.update(rowMap, sq.Eq{"id": user.Id})
	return user, err
}
