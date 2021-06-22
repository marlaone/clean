package apps

import (
	"github.com/go-chi/chi/v5"
	"marla.one/clean/pkg/database"
	"marla.one/clean/pkg/interfaces"
)

type AppsSystem struct {
	db     *database.DatabaseAdapater
	router *chi.Mux
}

func NewAppsSystem(db *database.DatabaseAdapater, router *chi.Mux) *AppsSystem {
	return &AppsSystem{
		db:     db,
		router: router,
	}
}

func (system *AppsSystem) Register(app interfaces.App) {

	if _, ok := app.(interfaces.DatabaseApp); ok {
		app.(interfaces.DatabaseApp).SetDatabaseAdapter(system.db)
	}

	app.OnCreate()

	if _, ok := app.(interfaces.HttpRoutableApp); ok {
		app.(interfaces.HttpRoutableApp).SetHttpRouter(system.router)
		app.(interfaces.HttpRoutableApp).PresentHttp()
	}

}
