package v1

import "net/http"

func (h *Handlers) GetTables(w http.ResponseWriter, r *http.Request) {
	h.usecase.GetTables(r.Context())

	w.Write([]byte("some tables"))
}
