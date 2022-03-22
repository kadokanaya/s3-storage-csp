package usecases_test

import (
	"sp/src/asserts"
	"sp/src/domains/entities"
	"sp/src/mocks"
	"sp/src/usecases/interactor"
	"testing"
)

func TestUserCreate(t *testing.T) {
	outputPort := mocks.NewUserOutputPortMock()
	repository := mocks.NewUserRepositoryMock()
	inputPort := interactor.NewUserInputPort(outputPort, repository)
	user := &entities.User{Address: "sdf", PubKey: []byte("sdf"), PrivKey: []byte("sdf")}

	created, err := inputPort.Create(user)
	if err != nil {
		t.Fatal(err)
	}
	asserts.AssertEqual(t, created.Address, user.Address, "アドレスが一致していません")
	asserts.AssertEqual(t, created.PubKey, user.PubKey, "公開鍵が一致していません")
	asserts.AssertEqual(t, created.PubKey, user.PubKey, "公開鍵が一致していません")
}

func TestUserFindByID(t *testing.T) {
	outputPort := mocks.NewUserOutputPortMock()
	repository := mocks.NewUserRepositoryMock()
	inputPort := interactor.NewUserInputPort(outputPort, repository)

	user := &entities.User{Address: "sdf", PubKey: []byte("sdf"), PrivKey: []byte("sdf")}

	id := "7"
	found, err := inputPort.FindByID(id)
	if err != nil {
		t.Fatal(err)
	}
	asserts.AssertEqual(t, found.ID, id, "IDが一致していません")
	asserts.AssertEqual(t, found.Address, user.Address, "アドレスが一致していません")
	asserts.AssertEqual(t, found.PubKey, user.PubKey, "公開鍵が一致していません")
	asserts.AssertEqual(t, found.PubKey, user.PubKey, "公開鍵が一致していません")
}
