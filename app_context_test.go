package clean_test

import (
	"net/http"
	"testing"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/interfaces"
)

type MockPresenter struct {
	interfaces.Registrable
	interfaces.AppContextable
}

var _ interfaces.Presenter = &MockPresenter{}

func NewMockPresenter(registry interfaces.Registry, ctx interfaces.AppContextable) *MockPresenter {
	return &MockPresenter{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: ctx,
	}
}

func (p *MockPresenter) CreateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *MockPresenter) ReadAction(w http.ResponseWriter, r *http.Request) {

}

func (p *MockPresenter) UpdateAction(w http.ResponseWriter, r *http.Request) {

}

func (p *MockPresenter) DeleteAction(w http.ResponseWriter, r *http.Request) {

}

func TestAppContextPresenter(t *testing.T) {

	registry := clean.NewRegistry()

	ctx := clean.NewAppContext(registry)

	ctx.RegisterPresenter("test", "mock", NewMockPresenter(registry, ctx))

	_, err := ctx.GetPresenter("test", "mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("presenter registration works")
}

type MockUseCase struct {
	interfaces.Registrable
	interfaces.AppContextable
}

var _ interfaces.UseCase = &MockUseCase{}

func NewMockUseCase(registry interfaces.Registry, ctx interfaces.AppContextable) *MockUseCase {
	return &MockUseCase{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: ctx,
	}
}

func TestAppContextUseCase(t *testing.T) {
	registry := clean.NewRegistry()

	ctx := clean.NewAppContext(registry)

	ctx.RegisterUseCase("mock", NewMockPresenter(registry, ctx))

	_, err := ctx.GetUseCase("mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("usecase registration works")
}

type MockRepository struct {
	interfaces.Registrable
	interfaces.AppContextable
}

var _ interfaces.Repository = &MockRepository{}

func NewMockRepository(registry interfaces.Registry, ctx interfaces.AppContextable) *MockRepository {
	return &MockRepository{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: ctx,
	}
}

func (repo *MockRepository) Create(e interfaces.Entity) (interfaces.Entity, error) {
	return e, nil
}

func (repo *MockRepository) Read(q interfaces.Query) ([]interfaces.Entity, error) {
	return []interfaces.Entity{}, nil
}

func (repo *MockRepository) Update(e interfaces.Entity) (interfaces.Entity, error) {
	return e, nil
}

func (repo *MockRepository) Delete(e interfaces.Entity) error {
	return nil
}

func TestAppContextRepository(t *testing.T) {
	registry := clean.NewRegistry()

	ctx := clean.NewAppContext(registry)

	ctx.RegisterRepository("mock", NewMockRepository(registry, ctx))

	_, err := ctx.GetRepository("mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("repository registration works")
}
