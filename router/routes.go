package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/handlers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/categories", handlers.Repo.GetCategories)
	mux.Post("/categories", handlers.Repo.CreateCategory)
	mux.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Check Page")
		if err != nil {
			fmt.Println(err)
		}
	})
	return mux
}
