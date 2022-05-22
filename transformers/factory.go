package transformers

import (
	"github.com/ibllex/go-fractal"
	"github.com/kachit/golang-api-skeleton/infrastructure"
	"github.com/kachit/golang-api-skeleton/models"
)

type Factory struct {
	fractal *fractal.Manager
	hashIds *infrastructure.HashIds
}

func NewTransformersFactory(fractal *fractal.Manager, hashIds *infrastructure.HashIds) *Factory {
	return &Factory{fractal, hashIds}
}

func (f *Factory) MapUsersResourceItem(user *models.User) (interface{}, error) {
	resource := f.NewUsersResourceItem(user)
	dataMap, err := f.fractal.CreateData(resource, nil).ToMap()
	if err != nil {
		return nil, err
	}
	return dataMap["data"], nil
}

func (f *Factory) MapUsersResourceCollection(users []*models.User) (interface{}, error) {
	collection := transformUsersToFractal(users)
	resource := f.NewUsersResourceCollection(collection)
	dataMap, err := f.fractal.CreateData(resource, nil).ToMap()
	if err != nil {
		return nil, err
	}
	return dataMap["data"], nil
}

func (f *Factory) NewUsersResourceItem(user *models.User) *fractal.Item {
	return fractal.NewItem(
		fractal.WithData(user),
		fractal.WithResourceKey("users"),
		fractal.WithTransformer(NewUsersTransformer(f.hashIds)),
	)
}

func (f *Factory) NewUsersResourceCollection(users []fractal.Any) *fractal.Collection {
	return fractal.NewCollection(
		fractal.WithData(users),
		fractal.WithResourceKey("users"),
		fractal.WithTransformer(NewUsersTransformer(f.hashIds)),
	)
}
