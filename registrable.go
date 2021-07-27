package clean

import "github.com/marlaone/clean/interfaces"

type Registrable struct {
	registry interfaces.Registry
}

var _ interfaces.Registrable = &Registrable{}

func NewRegistrable(registry interfaces.Registry) *Registrable {
	return &Registrable{
		registry: registry,
	}
}

func (a *Registrable) SetRegistry(r interfaces.Registry) {
	a.registry = r
}

func (a *Registrable) GetRegistry() interfaces.Registry {
	return a.registry
}
