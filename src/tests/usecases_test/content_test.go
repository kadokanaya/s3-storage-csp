package usecases_test

import (
	"sp/src/domains/entities"
	"sp/src/interfaces/storage"
	"sp/src/mocks"
	"sp/src/usecases/interactor"
	"testing"
)

func TestContentUpload(t *testing.T) {
	outputPort := mocks.NewContentOutputPortMock()
	repository := mocks.NewContentRepositoryMock()
	contract := mocks.NewContentContractMock()
	storage:= storage.NewContentStorage()
	userRepo := mocks.NewUserRepositoryMock()
	inputPort := interactor.NewContentInputPort(outputPort, repository, contract, storage, userRepo)

	contentInput := &entities.Content{
		Content:     []byte{1},
		MetaData:    [][]byte{},
		HashedData:  [][]byte{},
		ContentName: "コンテンツ1",
		SplitCount:  0,
		Owner:       "オーナー1",
		Id:          "12",
		UserId:      "1",
	}

	receipt, err := inputPort.Upload(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Id != contentInput.Id {
		t.Errorf("Content.Upload() should return entities.Receipt.Id = 7, but got = %s", receipt.Id)
	}
}
