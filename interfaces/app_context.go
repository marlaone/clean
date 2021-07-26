package interfaces

type AppContext interface {
	Registrable
	RegisterPresenter(presenterType string, name string, presenter Presenter)
	GetPresenter(presenterType string, name string) (Presenter, error)
	GetPresentersByType(presenterType string) (map[string]Presenter, error)
	RegisterUseCase(name string, uc UseCase)
	GetUseCase(name string) (UseCase, error)
	GetUseCases() map[string]UseCase
	RegisterRepository(name string, repo Repository)
	GetRepository(name string) (Repository, error)
	GetRepositories() map[string]Repository
}
