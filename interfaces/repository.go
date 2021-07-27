package interfaces

type Repository interface {
	Registrable
	AppContextable
	Create(Entity) (Entity, error)
	Read(Query) ([]Entity, error)
	Update(Entity) (Entity, error)
	Delete(Entity) error
}
