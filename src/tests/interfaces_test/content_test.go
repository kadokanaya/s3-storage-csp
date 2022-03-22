package interfaces_test

import (
	"sp/src/domains/entities"
	"sp/src/interfaces/storage"
	"sp/src/mocks"
	"testing"
)

// コンテンツ情報永続化モックのテスト
func TestContentRepositoryCreate(t *testing.T) {
	FakeRepo := mocks.NewContentRepositoryMock()
	testByte := []byte{1}
	contentInput := &entities.Content{
		Content:     []byte{1},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "コンテンツ1",
		SplitCount:  0,
		Owner:       "オーナー",
		Id:          "",
		UserId:      "12",
	}
	content, err := FakeRepo.Create(contentInput)
	if err != nil {
		t.Fatal(err)
	}
	if content.Id != contentInput.Id {
		t.Errorf("Content.Create() should return entities.Content.Id = %s, but got = %s", contentInput.Id, content.Id)
	}
}

// ブロックチェーン登録モックのテスト
func TestContentContractRegister(t *testing.T) {
	FakeContract := mocks.NewContentContractMock()
	testByte := []byte{1}
	content := &entities.Content{
		Content:     []byte{1},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
	}
	err := FakeContract.Register(content)
	if err != nil {
		t.Fatal(err)
	}
}

// ストレージ記録モックのテスト
func TestContentStorageCreate(t *testing.T) {
	cs := storage.NewContentStorage()
	testByte := []byte{1}
	content := &entities.Content{
		Content:     []byte{1},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "test1",
		SplitCount:  0,
		Owner:       "",
		Id:          "sdfsdf",
		UserId:      "",
		ContentId:   "dfg",
	}
	_, err := cs.Create(content)
	if err != nil {
		t.Fatal(err)
	}
}
