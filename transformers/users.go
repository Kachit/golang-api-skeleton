package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
)

func NewUsersTransformer(hashIds *infrastructure.HashIds) *UsersTransformer {
	return &UsersTransformer{&fractal.BaseTransformer{}, hashIds}
}

type UsersTransformer struct {
	*fractal.BaseTransformer
	hashIds *infrastructure.HashIds
}

func (t *UsersTransformer) Transform(data fractal.Any) fractal.M {
	result := fractal.M{}

	if u := t.toUser(data); u != nil {
		hash, _ := t.hashIds.EncodeUint64(u.Id)
		result["id"] = hash
		result["name"] = u.Name
		result["email"] = u.Email
		result["created_at"] = u.CreatedAt
		if u.ModifiedAt != nil && !u.ModifiedAt.IsZero() {
			result["modified_at"] = u.ModifiedAt
		}
		if !u.DeletedAt.Time.IsZero() {
			result["deleted_at"] = u.DeletedAt.Time
		}
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

func transformUsersToFractal(users []*models.User) []fractal.Any {
	var u []fractal.Any
	for _, user := range users {
		u = append(u, *user)
	}
	return u
}
