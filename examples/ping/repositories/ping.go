package repositories

import (
	"errors"
	"fmt"

	"github.com/marlaone/clean"
	"github.com/marlaone/clean/examples/ping/adapters"
	"github.com/marlaone/clean/interfaces"
)

type PingRepository struct {
	interfaces.Registrable
	interfaces.AppContextable
	adapter interfaces.StorageAdapter
}

var _ interfaces.Repository = &PingRepository{}

func NewPingRepository(registry interfaces.Registry, appContext interfaces.AppContextable) *PingRepository {
	adapter, err := registry.GetStorageAdapter("messages")

	if err != nil {
		panic(err)
	}

	return &PingRepository{
		Registrable:    clean.NewRegistrable(registry),
		AppContextable: appContext,
		adapter:        adapter,
	}
}

func (r *PingRepository) Read(q interfaces.Query) ([]interfaces.Entity, error) {
	adapter, err := r.GetRegistry().GetStorageAdapter("messages")

	if err != nil {
		return nil, err
	}

	messagesAdapter, ok := adapter.(*adapters.MessagesAdapter)

	if !ok {
		return nil, fmt.Errorf("missing messages adapter")
	}

	return []interfaces.Entity{
		messagesAdapter.Query("pong"),
	}, nil
}

func (r *PingRepository) Create(e interfaces.Entity) (interfaces.Entity, error) {
	return nil, errors.New("not implemented")
}

func (r *PingRepository) Update(e interfaces.Entity) (interfaces.Entity, error) {
	return nil, errors.New("not implemented")
}

func (r *PingRepository) Delete(e interfaces.Entity) error {
	return errors.New("not implemented")
}
