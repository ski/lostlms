// package main
//
// import (
//
//	"net/http"
//
//	"github.com/a-h/templ"
//	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/chi/v5/middleware"
//	"github.com/ski/lostlms/views"
//
// )
//
//	func main() {
//		r := chi.NewRouter()
//		r.Use(middleware.Logger)
//		component := views.Root("Hello Chi World")
//
//		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//			templ.Handler(component)
//		})
//		err := http.ListenAndServe(":3000", r)
//		if err != nil {
//			return
//		}
//	}
package main

import (
	"context"
	"github.com/ski/lostlms/config"
	"github.com/ski/lostlms/endpoints"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	server := endpoints.NewServer(cfg.HTTPServer)
	server.Start(ctx)
}
