package presenters

import (
	"encoding/json"
	"net/http"

	"sp/src/domains/entities"
	"sp/src/usecases/port"
)

type Audit struct {
	w http.ResponseWriter
}

func NewAuditOutputPort(w http.ResponseWriter) port.AuditOutputPort {
	return &Audit{
		w: w,
	}
}

func (u *Audit) Render(proofs *entities.Proofs, statusCode int) {
	res, err := json.Marshal(proofs)
	if err != nil {
		http.Error(u.w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.w.WriteHeader(statusCode)
	u.w.Header().Set("Audit-Type", "application/json")
	u.w.Write(res)
}

func (u *Audit) RenderError(err error, code int) {
	u.w.WriteHeader(code)
	http.Error(u.w, err.Error(), code)
}
