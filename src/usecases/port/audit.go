package port

import (
	entities "sp/src/domains/entities"
)

type AuditInputPort interface {
	Challen() (*entities.Proofs, error)
}

type AuditOutputPort interface {
	Render(*entities.Proofs, int)
	RenderError(error, int)
}

type AuditRepository interface {
	Create(proof *entities.Proof) (*entities.Proof, error)
	Update(proof *entities.Proof) (*entities.Proof, error)
}

type AuditContract interface {
	RegisterProof( content *entities.Proof ) error
	GetChallen(string) (*entities.Chal, error)
	GetContentLog(string) (*entities.Content, error)
}

type AuditStorage interface {
	GetContent(string) (*entities.Content, error)
}

type AuditCrypt interface {
	AuditProofGen( chal *entities.Chal, content *entities.Content, contentLog *entities.Content ) (*entities.Proof, error)
}