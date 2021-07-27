package interfaces

type App interface {
	Registrable
	SetContext(AppContext)
	GetContext() AppContext
	Setup()
	Run()
}
