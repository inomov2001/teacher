package sobject

import "github.com/google/uuid"

type Subject struct {
	ID          uuid.UUID
	Name        string
	Description string
}
