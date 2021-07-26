package main

import (
	"fmt"

	"github.com/marlaone/clean/interfaces"
)

type CleanAppContext struct {
	*Registrable
	presenters   map[string]map[string]interfaces.Presenter
	usecases     map[string]interfaces.UseCase
	repositories map[string]interfaces.Repository
}

var _ interfaces.AppContext = &CleanAppContext{}

func NewAppContext(registry interfaces.Registry) *CleanAppContext {
	return &CleanAppContext{
		Registrable:  NewRegistrable(registry),
		presenters:   make(map[string]map[string]interfaces.Presenter),
		usecases:     map[string]interfaces.UseCase{},
		repositories: map[string]interfaces.Repository{},
	}
}

func (ctx *CleanAppContext) GetPresentersByType(presenterType string) (map[string]interfaces.Presenter, error) {
	typePresenters, ok := ctx.presenters[presenterType]

	if !ok {
		return nil, fmt.Errorf("presenter type not registered: %s", presenterType)
	}

	return typePresenters, nil
}

func (ctx *CleanAppContext) GetPresenter(presenterType string, name string) (interfaces.Presenter, error) {
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

func (ctx *CleanAppContext) RegisterPresenter(presenterType string, name string, presenter interfaces.Presenter) {
	_, err := ctx.GetPresentersByType(presenterType)

	if err != nil {
		ctx.presenters[presenterType] = make(map[string]interfaces.Presenter)
	}

	ctx.presenters[presenterType][name] = presenter
}

func (ctx *CleanAppContext) RegisterUseCase(name string, usecase interfaces.UseCase) {
	ctx.usecases[name] = usecase
}

func (ctx *CleanAppContext) GetUseCase(name string) (interfaces.UseCase, error) {
	uc, ok := ctx.usecases[name]

	if !ok {
		return nil, fmt.Errorf("usecase doesn't exist: %s", name)
	}

	return uc, nil
}

func (ctx *CleanAppContext) GetUseCases() map[string]interfaces.UseCase {
	return ctx.usecases
}

func (ctx *CleanAppContext) RegisterRepository(name string, repo interfaces.Repository) {
	ctx.repositories[name] = repo
}

func (ctx *CleanAppContext) GetRepository(name string) (interfaces.Repository, error) {
	repo, ok := ctx.repositories[name]

	if !ok {
		return nil, fmt.Errorf("repository doesn't exist: %s", name)
	}

	return repo, nil
}

func (ctx *CleanAppContext) GetRepositories() map[string]interfaces.Repository {
	return ctx.repositories
}
