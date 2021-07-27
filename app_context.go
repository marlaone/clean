package clean

import (
	"fmt"

	"github.com/marlaone/clean/interfaces"
)

type AppContext struct {
	interfaces.Registrable
	*AppContextable
}

var _ interfaces.AppContext = &AppContext{}

func NewAppContext(registry interfaces.Registry) *AppContext {
	return &AppContext{
		Registrable:    NewRegistrable(registry),
		AppContextable: NewAppContextable(),
	}
}

func (ctx *AppContext) RegisterPresenter(presenterType string, name string, presenter interfaces.Presenter) {
	_, err := ctx.GetPresentersByType(presenterType)

	if err != nil {
		ctx.presenters[presenterType] = make(map[string]interfaces.Presenter)
	}

	ctx.presenters[presenterType][name] = presenter
}

func (ctx *AppContext) RegisterUseCase(name string, usecase interfaces.UseCase) {
	ctx.usecases[name] = usecase
}

func (ctx *AppContext) GetUseCase(name string) (interfaces.UseCase, error) {
	uc, ok := ctx.usecases[name]

	if !ok {
		return nil, fmt.Errorf("usecase doesn't exist: %s", name)
	}

	return uc, nil
}

func (ctx *AppContext) RegisterRepository(name string, repo interfaces.Repository) {
	ctx.repositories[name] = repo
}
