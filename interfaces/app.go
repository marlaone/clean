package interfaces

import "context"

type App interface {
	Registrable
	SetContext(AppContext)
	GetContext() AppContext
	Setup()
	Run(context.Context)
}
