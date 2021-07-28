package clean

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/marlaone/clean/interfaces"
)

type Registry struct {
	apps     map[string]interfaces.App
	adapters map[string]interfaces.StorageAdapter
	eventBus EventBus.Bus
}

var _ interfaces.Registry = &Registry{}

func NewRegistry() *Registry {
	return &Registry{
		apps:     map[string]interfaces.App{},
		adapters: map[string]interfaces.StorageAdapter{},
		eventBus: EventBus.New(),
	}
}

func (r *Registry) RegisterApp(appName string, app interfaces.App) {
	r.apps[appName] = app
}

func (r *Registry) GetApp(appName string) (interfaces.App, error) {
	app, ok := r.apps[appName]

	if !ok {
		return nil, fmt.Errorf("app doesn't exist: %s", appName)
	}

	return app, nil
}

func (r *Registry) GetApps() map[string]interfaces.App {
	return r.apps
}

func (r *Registry) RegisterStorageAdapter(name string, adapter interfaces.StorageAdapter) {
	r.adapters[name] = adapter
}

func (r *Registry) GetStorageAdapter(name string) (interfaces.StorageAdapter, error) {
	adapter, ok := r.adapters[name]

	if !ok {
		return nil, fmt.Errorf("adapter doesn't exist: %s", name)
	}

	return adapter, nil
}

func (r *Registry) GetEventBus() EventBus.Bus {
	return r.eventBus
}

func (r *Registry) SetEventBus(eb EventBus.Bus) {
	r.eventBus = eb
}
