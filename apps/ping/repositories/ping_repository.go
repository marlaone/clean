package repositories

import (
	"marla.one/clean/apps/ping/entities"
	"marla.one/clean/pkg/interfaces"
)

type PingRepository struct {
}

var _ interfaces.Repository = &PingRepository{}

func NewPingRepository() *PingRepository {
	return &PingRepository{}
}

func (repo *PingRepository) Fetch(filter interfaces.Filter) ([]interfaces.Entity, error) {
	return []interfaces.Entity{
		&entities.PingMessage{Message: "pong"},
	}, nil
}

func (repo *PingRepository) FetchOne(filter interfaces.Filter) (interfaces.Entity, error) {
	return &entities.PingMessage{Message: "pong"}, nil
}

func (repo *PingRepository) Save(message interfaces.Entity) (interfaces.Entity, error) {
	return message, nil
}

func (repo *PingRepository) Delete(message interfaces.Entity) error {
	return nil
}
