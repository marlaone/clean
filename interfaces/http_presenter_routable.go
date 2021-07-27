package interfaces

type HttpPresenterRoutable interface {
	SetRoutes(map[string]string) error
	GetRoutes() (createRoute string, readRoute string, updateRoute string, deleteRoute string)
}
