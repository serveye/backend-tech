package backend_tech

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/controllers"
	"github.com/serveye/backend-tech/render"
	"log"
	"net/http"
)

const portNumber = ":8080"



func Execute() {

	router := mux.NewRouter()
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")

	//router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/check", func (w http.ResponseWriter, r *http.Request) {
		_, err :=fmt.Fprint(w, "Check Page")
		if err != nil {
			fmt.Println(err)
		}
	}).Methods("GET")
	var appConfig config.AppConfig
	  tc, err := render.CreateTemplateCache()
	  if err != nil {
	  	log.Fatal("Cannot create template cache")
	  }

	  appConfig.TemplateCache = tc

	  render.NewTemplateCache(&appConfig)
	fmt.Println(
		fmt.Sprintf("Starting application on port %s", portNumber),
	)

	err = http.ListenAndServe(portNumber, router)
	if err != nil {
		fmt.Println("server failed to start", err)
	}
}

