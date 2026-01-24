package http

import (
	"ca-example/internal/usecase"
	"net/http"
)

type MemberHandler struct {
	UC *usecase.RewardPointsUseCase
}

func (h *MemberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.UC.Execute(id, 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
