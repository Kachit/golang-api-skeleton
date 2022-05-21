package transformers

import (
	"github.com/ibllex/go-fractal"
)

// DataArraySerializer array serializer with default resource key
type DataArraySerializer struct {
	fractal.ArraySerializer
}

// Collection serialize a collection
func (s *DataArraySerializer) Collection(resourceKey string, data fractal.Any) fractal.M {
	return data.(fractal.M)
}

// Item serialize an item
func (s *DataArraySerializer) Item(resourceKey string, data fractal.Any) fractal.M {
	return data.(fractal.M)
}

func NewFractalManager() *fractal.Manager {
	manager := fractal.NewManager(nil)
	manager.SetSerializer(&DataArraySerializer{})
	return manager
}
