package interfaces

import (
	"github.com/go-chi/chi/v5"
	"marla.one/clean/pkg/database"
)

type App interface {
	SetContext(AppContext)
	GetContext() AppContext
	OnCreate()
}

type AppContext interface {
	RegisterUseCase(name string, usecase UseCase)
	GetUseCase(name string) (UseCase, error)
	RegisterRepository(name string, repository Repository)
	GetRepository(name string) (Repository, error)
}

type DatabaseApp interface {
	SetDatabaseAdapter(*database.DatabaseAdapater)
	GetDatabaseAdapter() *database.DatabaseAdapater
}

type HttpRoutableApp interface {
	SetHttpRouter(*chi.Mux)
	GetHttpRouter() *chi.Mux
	PresentHttp()
	MountHttpRoutes() *chi.Mux
}

type UseCase interface {
}

type Repository interface {
	Fetch(Filter) ([]Entity, error)
	FetchOne(Filter) (Entity, error)
	Save(Entity) (Entity, error)
	Delete(Entity) error
}

type Filter interface {
}

type Entity interface {
}
