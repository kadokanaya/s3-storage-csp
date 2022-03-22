package interfaces_test

import (
	"sp/src/domains/entities"
	"sp/src/mocks"
	"testing"
)

// コンテンツ情報永続化モックのテスト
func TestUserRepositoryCreate(t *testing.T) {
	FakeRepo := mocks.NewUserRepositoryMock()
	userInput := &entities.User{
		Address: "sdf",
		PubKey:  []byte("sdf"),
		PrivKey: []byte("sdf"),
	}
	user, err := FakeRepo.Create(userInput)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != "7" {
		t.Errorf("User.Create() should return entities.User.ID = %s, but got = %s", userInput.ID, user.ID)
	}
}
