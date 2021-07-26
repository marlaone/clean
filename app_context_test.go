package main

import (
	"testing"

	"github.com/marlaone/clean/interfaces"
)

type MockPresenter struct {
	*Registrable
}

var _ interfaces.Presenter = &MockPresenter{}

func NewMockPresenter(registry interfaces.Registry) *MockPresenter {
	return &MockPresenter{
		Registrable: NewRegistrable(registry),
	}
}

func TestAppContextPresenter(t *testing.T) {

	registry := NewCleanRegistry()

	ctx := NewAppContext(registry)

	ctx.RegisterPresenter("test", "mock", NewMockPresenter(registry))

	_, err := ctx.GetPresenter("test", "mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("presenter registration works")
}

type MockUseCase struct {
	*Registrable
}

var _ interfaces.UseCase = &MockUseCase{}

func NewMockUseCase(registry interfaces.Registry) *MockUseCase {
	return &MockUseCase{
		Registrable: NewRegistrable(registry),
	}
}

func TestAppContextUseCase(t *testing.T) {
	registry := NewCleanRegistry()

	ctx := NewAppContext(registry)

	ctx.RegisterUseCase("mock", NewMockPresenter(registry))

	_, err := ctx.GetUseCase("mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("usecase registration works")
}

type MockRepository struct {
	*Registrable
}

var _ interfaces.Repository = &MockRepository{}

func NewMockRepository(registry interfaces.Registry) *MockRepository {
	return &MockRepository{
		Registrable: NewRegistrable(registry),
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
	registry := NewCleanRegistry()

	ctx := NewAppContext(registry)

	ctx.RegisterRepository("mock", NewMockRepository(registry))

	_, err := ctx.GetRepository("mock")

	if err != nil {
		t.Error(err.Error())
	}

	t.Log("repository registration works")
}
