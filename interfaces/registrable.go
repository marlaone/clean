package interfaces

type Registrable interface {
	SetRegistry(Registry)
	GetRegistry() Registry
}
