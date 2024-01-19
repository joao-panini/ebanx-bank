package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	OK bool
}

func (h *handler) ResetHandler(w http.ResponseWriter, r *http.Request) {

	err := h.accService.ResetAccountStates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(0)
	}

	w.Header().Set(ContentType, PlainTextContentType)
	w.Write([]byte("OK"))

}
