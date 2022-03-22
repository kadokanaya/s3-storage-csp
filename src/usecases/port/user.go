package port

import (
	entities "sp/src/domains/entities"
)

type UserInputPort interface {
	Create(*entities.User) (*entities.User, error)
	FindByID(string) (*entities.User, error)
}

type UserOutputPort interface {
	Render(*entities.User, int)
	RenderError(error, int)
}

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
	FindByID(string) (*entities.User, error)
	FindByAddress(string) (*entities.User, error)
}
