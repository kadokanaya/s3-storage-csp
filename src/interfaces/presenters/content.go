package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sp/src/domains/entities"
	"sp/src/usecases/port"
)

type Content struct {
	w http.ResponseWriter
}

func NewContentOutputPort(w http.ResponseWriter) port.ContentOutputPort {
	return &Content{
		w: w,
	}
}

func (u *Content) Render(Content *entities.Receipt, statusCode int) {
	fmt.Println(Content)
	res, err := json.Marshal(Content)
	if err != nil {
		http.Error(u.w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.w.WriteHeader(statusCode)
	u.w.Header().Set("Content-Type", "application/json")
	u.w.Write(res)
}

func (u *Content) RenderURL(url string, code int) {
	res, err := json.Marshal(url)
	if err != nil {
		http.Error(u.w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.w.WriteHeader(code)
	u.w.Header().Set("Content-Type", "application/json")
	u.w.Write(res)
}

func (u *Content) RenderError(err error, code int) {
	u.w.WriteHeader(code)
	http.Error(u.w, err.Error(), code)
}
