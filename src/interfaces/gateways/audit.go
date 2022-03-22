package gateways

import (
	"sp/src/domains/entities"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ProofSQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.Proof, error)
	First(interface{}, ...interface{}) (*entities.Proof, error)
	Create(interface{}) (*entities.Proof, error)
	Save(interface{}) (*entities.Proof, error)
	Delete(interface{}) *entities.Proof
	Where(interface{}, ...interface{}) *entities.Proof
}

type proofRepository struct {
	Conn *gorm.DB
	ProofSQLHandler
}

func NewProofRepository(conn *gorm.DB) port.AuditRepository {
	return &proofRepository{
		Conn: conn,
	}
}

func (pr *proofRepository) Create(u *entities.Proof) (*entities.Proof, error) {
	created, err := pr.ProofSQLHandler.Create(u)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (pr *proofRepository) Update(u *entities.Proof) (*entities.Proof, error) {
	updated, err := pr.ProofSQLHandler.Save(u)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
