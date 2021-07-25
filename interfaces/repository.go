package interfaces

type Repository interface {
	Create(Entity) (Entity, error)
	Read(Query) ([]Entity, error)
	Update(Entity) (Entity, error)
	Delete(Entity) error
}
