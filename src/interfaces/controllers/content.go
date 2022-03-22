package controllers

import (
	"encoding/json"
	"net/http"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/interfaces/gateways"
	"sp/src/interfaces/presenters"
	"sp/src/interfaces/storage"
	"sp/src/mocks"
	"sp/src/usecases/interactor"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func() port.ContentContract
	// -> crypt.NewContentCrypt
	CryptFactory func() port.ContentContract
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		co port.ContentContract,
	) port.ContentInputPort
	Conn *gorm.DB
}

func LoadContentController(db *gorm.DB) *ContentController {
	return &ContentController{Conn: db}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	content := &entities.Content{}
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputPort := presenters.NewContentOutputPort(w)
	repository := gateways.NewContentRepository(cc.Conn)
	contract := mocks.NewContentContractMock()
	storage := storage.NewContentStorage()
	userRepo := gateways.NewUserRepository(cc.Conn)
	inputPort := interactor.NewContentInputPort(
		outputPort,
		repository,
		contract,
		storage,
		userRepo,
	)
	inputPort.Upload(content)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	_, tail := core.ShiftPath(r.URL.Path)
	_, tail = core.ShiftPath(tail)
	id, _ := core.ShiftPath(tail)
	outputPort := presenters.NewContentOutputPort(w)
	repository := gateways.NewContentRepository(cc.Conn)
	contract := mocks.NewContentContractMock()
	storage := storage.NewContentStorage()
	userRepo := gateways.NewUserRepository(cc.Conn)
	inputPort := interactor.NewContentInputPort(
		outputPort,
		repository,
		contract,
		storage,
		userRepo,
	)
	inputPort.FindByID(id)
}

func (cc *ContentController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cc.Post(w, r)
	case "GET":
		cc.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}
