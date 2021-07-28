package interfaces

import "github.com/asaskevich/EventBus"

type Registry interface {
	RegisterApp(appName string, app App)
	GetApp(appName string) (App, error)
	GetApps() map[string]App
	RegisterStorageAdapter(name string, adapter StorageAdapter)
	GetStorageAdapter(name string) (StorageAdapter, error)
	SetEventBus(EventBus.Bus)
	GetEventBus() EventBus.Bus
}
