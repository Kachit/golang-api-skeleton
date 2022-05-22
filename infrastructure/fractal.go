package infrastructure

import (
	"github.com/ibllex/go-fractal"
)

func NewFractalManager() *fractal.Manager {
	manager := fractal.NewManager(nil)
	manager.SetSerializer(&fractal.DataArraySerializer{})
	return manager
}
