package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/models"
)

func MapUsersResourceItem(user *models.User) (fractal.M, error) {
	resource := NewUsersResourceItem(user)
	return NewFractalManager().CreateData(resource, nil).ToMap()
}

func MapUsersResourceCollection(users []*models.User) (fractal.M, error) {
	resource := NewUsersResourceCollection(users)
	return NewFractalManager().CreateData(resource, nil).ToMap()
}

func NewUsersResourceItem(user *models.User) *fractal.Item {
	return fractal.NewItem(
		fractal.WithData(user),
		fractal.WithResourceKey("users"),
		fractal.WithTransformer(NewUsersTransformer()),
	)
}

func NewUsersResourceCollection(users []*models.User) *fractal.Collection {
	return fractal.NewCollection(
		fractal.WithData(users),
		fractal.WithResourceKey("users"),
		fractal.WithTransformer(NewUsersTransformer()),
	)
}

func NewUsersTransformer() *UsersTransformer {
	return &UsersTransformer{&fractal.BaseTransformer{}}
}

type UsersTransformer struct {
	*fractal.BaseTransformer
}

func (t *UsersTransformer) Transform(data fractal.Any) fractal.M {
	result := fractal.M{}

	if u := t.toUser(data); u != nil {
		result["id"] = u.Id
		result["name"] = u.Name
		result["email"] = u.Email
		result["created_at"] = u.CreatedAt
	}
	return result
}

func (t *UsersTransformer) toUser(data fractal.Any) *models.User {
	switch b := data.(type) {
	case *models.User:
		return b
	case models.User:
		return &b
	}
	return nil
}
