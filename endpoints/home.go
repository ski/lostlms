package endpoints

import (
	"github.com/ski/lostlms/views"
	"net/http"
)

func (s *Server) getHome(w http.ResponseWriter, r *http.Request) {
	c := views.Root("World")
	err := c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	return
}

func (s *Server) postClicked(w http.ResponseWriter, r *http.Request) {
	c := views.Clicked("World")
	err := c.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	return
}
