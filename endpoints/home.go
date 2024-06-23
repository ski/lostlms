package endpoints

import (
	"github.com/ski/lostlms/views"
	"net/http"
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	c := views.Root("Cock Sucker")
	err := c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	return
}
