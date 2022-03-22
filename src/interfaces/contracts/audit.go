package contracts

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"
)

type AuditContract struct{}

func NewAuditContracts() port.AuditContract {
	return &AuditContract{}
}

func (cc *AuditContract) GetChallen(id string) (*entities.Chal, error) {
	return &entities.Chal{}, nil
}

func (cc *AuditContract) GetContentLog(id string) (*entities.Content, error) {
	return &entities.Content{}, nil
}

func (cc *AuditContract) RegisterProof(proof *entities.Proof) error {
	return nil
}
