package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sp/src/asserts"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/controllers"
	"sp/src/interfaces/gateways"
	"sp/src/interfaces/presenters"
	"sp/src/interfaces/storage"
	"sp/src/usecases/interactor"
	"testing"
)

//* 登録済みユーザのuserIdでコンテンツを作成
func TestCreateContent(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	ur := gateways.NewUserRepository(db)
	u := &entities.User{
		ID:      "sdsdsd",
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	user, _ := ur.Create(u)
	var buf bytes.Buffer
	testByte := []byte{1}
	if err := json.NewEncoder(&buf).Encode(&entities.Content{
		Content:     []byte{1, 2, 3, 4, 5},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "コンテンツ1",
		SplitCount:  2,
		Owner:       "sdf",
		Id:          "sdf",
		UserId:      user.ID,
		ContentId:   "sdf",
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/content", &buf)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(db)
	cc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Result().Status)
}

//* 登録済みユーザのuserIdでコンテンツを取得
func TestGetContent(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	//* ユーザ作成
	u := &entities.User{
		Address: "sdf",
		PubKey:  []byte("pubKey"),
		PrivKey: []byte("privKey"),
	}
	recCreate := httptest.NewRecorder()
	uo := presenters.NewUserOutputPort(recCreate)
	ur := gateways.NewUserRepository(db)
	ui := interactor.NewUserInputPort(uo, ur)
	user, err := ui.Create(u)
	if err != nil {
		t.Errorf("Failed to Create User: %v", err)
	}
	testByte := []byte{1}
	//* 作成したユーザのIDでコンテンツをアップロード
	c := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "sdf",
		UserId:      user.ID,
	}
	recUpload := httptest.NewRecorder()
	co := presenters.NewContentOutputPort(recUpload)
	cr := gateways.NewContentRepository(db)
	cs := storage.NewContentStorage()
	cco := contracts.NewContentContracts()
	ci := interactor.NewContentInputPort(co, cr, cco, cs, ur)
	content, err := ci.Upload(c)
	if err != nil {
		t.Errorf("Failed to Create Content: %v", err)
	}
	//*　作成したコンテンツのIDを用いてコンテンツを取得
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/contents/%s", content.Id), nil)
	rec := httptest.NewRecorder()
	cc := controllers.LoadContentController(db)
	cc.Dispatch(rec, req)
	receipt := &entities.Receipt{}
	err = json.NewDecoder(rec.Body).Decode(&receipt)
	if err != nil {
		t.Errorf("Failed to Get Content: %v", err)
	}
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
	asserts.AssertEqual(t, content, receipt, rec.Result().Status)
}

//* 存在しないUserIDで作成する -> エラー
func TestGetContentError(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/contents/%s", "sdfsdf"), nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadUserController(db)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusBadRequest, rec.Code, rec.Result().Status)
}
