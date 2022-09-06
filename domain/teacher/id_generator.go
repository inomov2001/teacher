package teacher

import "github.com/google/uuid"

// IDGenerator is interface that abstract away actual way of generating new uuids
type IDGenerator interface {
	GeneratorID() uuid.UUID
}
type uuidGenerator struct{}

func (g uuidGenerator) GeneratorID() uuid.UUID {
	return uuid.New()
}
