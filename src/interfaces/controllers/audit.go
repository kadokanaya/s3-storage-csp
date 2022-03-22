package controllers

import (
	"net/http"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"sp/src/interfaces/gateways"
	"sp/src/interfaces/presenters"
	"sp/src/interfaces/storage"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type AuditController struct {
	// -> gateway.NewAuditRepository
	RepoFactory func(c *gorm.DB) port.AuditRepository
	// -> contracts.NewAuditContracts
	ContractFactory func() port.AuditContract
	// -> presenter.NewAuditOutputPort
	OutputFactory func(w http.ResponseWriter) port.AuditOutputPort
	// -> crypt.NewAuditCrypt
	CryptFactory func(p *entities.Param) port.AuditCrypt
	// -> interactor.NewAuditInputPort
	InputFactory func(
		o port.AuditOutputPort,
		u port.AuditRepository,
		co port.AuditContract,
		cr port.AuditCrypt,
	) port.AuditInputPort
	Conn  *gorm.DB
	Param *entities.Param
}

func LoadAuditController(db *gorm.DB, param *entities.Param) *AuditController {
	return &AuditController{Conn: db, Param: param}
}

func (ac *AuditController) Post(w http.ResponseWriter, r *http.Request) {
	outputPort := presenters.NewAuditOutputPort(w)
	repository := gateways.NewProofRepository(ac.Conn)
	contract := contracts.NewAuditContracts()
	crypt := crypt.NewAuditCrypt(ac.Param)
	storage := storage.NewContentStorage()
	inputPort := interactor.NewAuditInputPort(outputPort, contract, crypt, storage, repository)
	inputPort.Challen()
}

func (ac *AuditController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ac.Post(w, r)
	default:
		http.NotFound(w, r)
	}
}
