package mocks

import (
	entities "sp/src/domains/entities"
	"sp/src/usecases/port"
)

type ContentControllerMock struct {
	RepoFactory     func() port.ContentRepository
	ContractFactory func() port.ContentContract
	OutputFactory   func() port.ContentOutputPort
	InputFactory    func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		cc port.ContentContract,
	) port.ContentInputPort
}

func NewContentRepositoryMock() port.ContentRepository {
	return &ContentRepositoryMock{}
}

func NewContentContractMock() port.ContentContract {
	return &ContentContractMock{}
}

func NewContentOutputPortMock() port.ContentOutputPort {
	return &ContentOutputPortMock{}
}

type ContentRepositoryMock struct {
}
type ContentOutputPortMock struct {
}

type ContentSPMock struct {
}

type ContentContractMock struct {
}

func (m *ContentRepositoryMock) Create(content *entities.Content) (*entities.Receipt, error) {
	created := &entities.Receipt{
		Id:           content.Id,
		ContentLogId: content.Id,
		ContentURL:   "localhost:3000/api/content/" + content.Id,
		FileName:     "asd",
	}
	return created, nil
}

func (m *ContentRepositoryMock) Find(id string) (*entities.Receipt, error) {
	found := &entities.Receipt{
		Id:           id,
		ContentLogId: id,
		ContentURL:   "localhost:3000/api/content/" + id,
		FileName:     "sdf",
	}
	return found, nil
}

func (m *ContentOutputPortMock) Render(*entities.Receipt, int) {
}

func (m *ContentOutputPortMock) RenderURL(string, int) {
}

func (m *ContentOutputPortMock) RenderError(error, int) {
}

func (m *ContentContractMock) Register(*entities.Content) error {
	return nil
}
