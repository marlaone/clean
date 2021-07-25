package main

import (
	"fmt"

	"github.com/marlaone/clean/interfaces"
)

type CleanRegistry struct {
	apps     map[string]interfaces.App
	adapters map[string]interfaces.StorageAdapter
}

var _ interfaces.Registry = &CleanRegistry{}

func NewCleanRegistry() *CleanRegistry {
	return &CleanRegistry{
		apps:     map[string]interfaces.App{},
		adapters: map[string]interfaces.StorageAdapter{},
	}
}

func (r *CleanRegistry) RegisterApp(appName string, app interfaces.App) {
	r.apps[appName] = app
}

func (r *CleanRegistry) GetApp(appName string) (interfaces.App, error) {
	app, ok := r.apps[appName]

	if !ok {
		return nil, fmt.Errorf("app doesn't exist: %s", appName)
	}

	return app, nil
}

func (r *CleanRegistry) RegisterStorageAdapter(name string, adapter interfaces.StorageAdapter) {
	r.adapters[name] = adapter
}

func (r *CleanRegistry) GetStorageAdapter(name string) (interfaces.StorageAdapter, error) {
	adapter, ok := r.adapters[name]

	if !ok {
		return nil, fmt.Errorf("adapter doesn't exist: %s", name)
	}

	return adapter, nil
}
