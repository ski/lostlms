package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

type healthResponse struct {
	OK bool `json:"ok"`
}

func (s *Server) handleGetHealth(w http.ResponseWriter, r *http.Request) {
	health := healthResponse{OK: true}
	err := render.Render(w, r, health)
	if err != nil {
		return
	}
}
