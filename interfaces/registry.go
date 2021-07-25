package interfaces

type Registry interface {
	RegisterApp(appName string, app App)
	GetApp(appName string) (App, error)
	RegisterStorageAdapter(name string, adapter StorageAdapter)
	GetStorageAdapter(name string) (StorageAdapter, error)
}
