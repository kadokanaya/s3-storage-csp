package mocks

import (
	entities "sp/src/domains/entities"
	"sp/src/usecases/port"
)

type UserControllerMock struct {
	RepoFactory   func() port.UserRepository
	OutputFactory func() port.UserOutputPort
	InputFactory  func(
		o port.UserOutputPort,
		u port.UserRepository,
	) port.UserInputPort
}

func NewUserRepositoryMock() port.UserRepository {
	return &userRepositoryMock{}
}

func NewUserOutputPortMock() port.UserOutputPort {
	return &userOutputPortMock{}
}

type userRepositoryMock struct {
}

type userOutputPortMock struct {
}

func (m *userRepositoryMock) Create(user *entities.User) (*entities.User, error) {
	created := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
	return created, nil
}

func (m *userRepositoryMock) FindByID(id string) (*entities.User, error) {
	user := &entities.User{ID: id, Address: "sdf", PubKey: []byte("sdf"), PrivKey: []byte("sdf")}
	return user, nil
}

func (m *userRepositoryMock) FindByAddress(address string) (*entities.User, error) {
	user := &entities.User{ID: "7", Address: address, PubKey: []byte("sdf"), PrivKey: []byte("sdf")}
	return user, nil
}


func (m *userRepositoryMock) Update(user *entities.User) (*entities.User, error) {
	updated := &entities.User{ID: "7", Address: user.Address, PubKey: user.PubKey, PrivKey: user.PrivKey}
	return updated, nil
}

func (m *userOutputPortMock) Render(*entities.User, int) {
}

func (m *userOutputPortMock) RenderError(error, int) {
}
