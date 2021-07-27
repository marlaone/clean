package clean

import (
	"fmt"

	"github.com/marlaone/clean/interfaces"
)

type AppContextable struct {
	presenters   map[string]map[string]interfaces.Presenter
	usecases     map[string]interfaces.UseCase
	repositories map[string]interfaces.Repository
}

var _ interfaces.AppContextable = &AppContext{}

func NewAppContextable() *AppContextable {
	return &AppContextable{
		presenters:   make(map[string]map[string]interfaces.Presenter),
		usecases:     map[string]interfaces.UseCase{},
		repositories: map[string]interfaces.Repository{},
	}
}

func (ctx *AppContextable) GetPresentersByType(presenterType string) (map[string]interfaces.Presenter, error) {
	typePresenters, ok := ctx.presenters[presenterType]

	if !ok {
		return nil, fmt.Errorf("presenter type not registered: %s", presenterType)
	}

	return typePresenters, nil
}

func (ctx *AppContextable) GetPresenter(presenterType string, name string) (interfaces.Presenter, error) {
	typePresenters, err := ctx.GetPresentersByType(presenterType)

	if err != nil {
		return nil, err
	}

	presenter, ok := typePresenters[name]

	if !ok {
		return nil, fmt.Errorf("presenter '%s' doesn't exist on type: %s", name, presenterType)
	}

	return presenter, nil
}

func (ctx *AppContextable) GetUseCase(name string) (interfaces.UseCase, error) {
	uc, ok := ctx.usecases[name]

	if !ok {
		return nil, fmt.Errorf("usecase doesn't exist: %s", name)
	}

	return uc, nil
}

func (ctx *AppContextable) GetUseCases() map[string]interfaces.UseCase {
	return ctx.usecases
}

func (ctx *AppContextable) GetRepository(name string) (interfaces.Repository, error) {
	repo, ok := ctx.repositories[name]

	if !ok {
		return nil, fmt.Errorf("repository doesn't exist: %s", name)
	}

	return repo, nil
}

func (ctx *AppContextable) GetRepositories() map[string]interfaces.Repository {
	return ctx.repositories
}
