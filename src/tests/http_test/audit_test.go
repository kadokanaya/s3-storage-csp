package http

import (
	"net/http"
	"net/http/httptest"
	"sp/src/asserts"
	"sp/src/interfaces/controllers"
	"testing"
)

// UserName, EmailのあるユーザをPOST -> 201を返すかをテスト
func TestCreateAudit(t *testing.T) {
	db, err := LoadTestDB()
	if err != nil {
		t.Errorf("Failed to get DB: %v", err)
		return
	}
	param, err := LoadTestParam()
	if err != nil {
		t.Errorf("Failed to get Param: %v", err)
		return
	}
	req := httptest.NewRequest(http.MethodPost, "/api/audit", nil)
	rec := httptest.NewRecorder()
	uc := controllers.LoadAuditController(db, param)
	uc.Dispatch(rec, req)
	asserts.AssertEqual(t, http.StatusOK, rec.Code, rec.Result().Status)
}
