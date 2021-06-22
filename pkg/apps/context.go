package apps

import (
	"fmt"

	"marla.one/clean/pkg/interfaces"
)

type AppContext struct {
	appName      string
	usecases     map[string]interfaces.UseCase
	repositories map[string]interfaces.Repository
}

var _ interfaces.AppContext = &AppContext{}

func NewAppContext(appName string) *AppContext {
	return &AppContext{
		appName:      appName,
		usecases:     map[string]interfaces.UseCase{},
		repositories: map[string]interfaces.Repository{},
	}
}

func (ctx *AppContext) RegisterUseCase(name string, usecase interfaces.UseCase) {
	ctx.usecases[name] = usecase
}

func (ctx *AppContext) GetUseCase(name string) (interfaces.UseCase, error) {

	if usecase, ok := ctx.usecases[name]; ok {
		return usecase, nil
	}

	return nil, fmt.Errorf("Invalid usecase %s for app %s", name, ctx.appName)
}

func (ctx *AppContext) RegisterRepository(name string, repository interfaces.Repository) {
	ctx.repositories[name] = repository
}

func (ctx *AppContext) GetRepository(name string) (interfaces.Repository, error) {
	if usecase, ok := ctx.repositories[name]; ok {
		return usecase, nil
	}

	return nil, fmt.Errorf("Invalid repository %s for app %s", name, ctx.appName)
}
