package api

import (
	"encoding/json"
	"net/http"
)

func (h *handler) ResetHandler(w http.ResponseWriter, r *http.Request) {
	err := h.accService.ResetAccountStates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(0)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}
