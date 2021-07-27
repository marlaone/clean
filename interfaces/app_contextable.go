package interfaces

type AppContextable interface {
	GetPresenter(presenterType string, name string) (Presenter, error)
	GetPresentersByType(presenterType string) (map[string]Presenter, error)
	GetUseCase(name string) (UseCase, error)
	GetUseCases() map[string]UseCase
	GetRepository(name string) (Repository, error)
	GetRepositories() map[string]Repository
}
