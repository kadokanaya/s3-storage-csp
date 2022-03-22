package mocks

import (
	entities "sp/src/domains/entities"
	"sp/src/usecases/port"
)

type AuditControllerMock struct {
	RepoFactory     func() port.AuditRepository
	ContractFactory func() port.AuditContract
	OutputFactory   func() port.AuditOutputPort
	CryptFactory    func(param entities.Param) port.AuditCrypt
	InputFactory    func(
		o port.AuditOutputPort,
		r port.AuditRepository,
		cc port.AuditContract,
		cr port.AuditCrypt,
	) port.AuditInputPort
}

func NewAuditRepositoryMock() port.AuditRepository {
	return &AuditRepositoryMock{}
}

func NewAuditContractMock() port.AuditContract {
	return &AuditContractMock{}
}

func NewAuditOutputPortMock() port.AuditOutputPort {
	return &AuditOutputPortMock{}
}

func NewAuditCryptMock(p entities.Param) port.AuditCrypt {
	return &AuditCryptMock{
		Param: p,
	}
}

type AuditRepositoryMock struct {
}

type AuditOutputPortMock struct {
}

type AuditCryptMock struct {
	Param entities.Param
}

type AuditContractMock struct {
}

func (m *AuditRepositoryMock) Create(proof *entities.Proof) (*entities.Proof, error) {
	created := &entities.Proof{
		Myu:       []byte{},
		Gamma:     []byte{},
		ContentId: "1",
	}
	return created, nil
}

func (m *AuditRepositoryMock) Update(proof *entities.Proof) (*entities.Proof, error) {
	updated := &entities.Proof{
		Myu:       []byte{},
		Gamma:     []byte{},
		ContentId: "1",
	}
	return updated, nil
}

func (m *AuditOutputPortMock) Render(*entities.Proofs, int) {
}

func (m *AuditOutputPortMock) RenderError(error, int) {
}

func (m *AuditContractMock) RegisterProof(*entities.Proof) error {
	return nil
}

func (m *AuditContractMock) GetChallen(id string) (*entities.Chal, error) {
	chal := &entities.Chal{
		ContentId: id,
		C:         0,
		K1:        []byte{},
		K2:        []byte{},
	}
	return chal, nil
}

func (m *AuditCryptMock) AuditProofGen(chal *entities.Chal, content *entities.Content, contentLog *entities.Content) (*entities.Proof, error) {
	proof := &entities.Proof{
		Myu:       []byte{1},
		Gamma:     []byte{1},
		ContentId: chal.ContentId,
	}
	return proof, nil
}

func (m *AuditContractMock) GetContentLog(id string) (*entities.Content, error) {
	content := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          id,
		UserId:      id,
		ContentId:   "sdf",
	}
	return content, nil
}