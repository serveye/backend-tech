package backend_tech

import (
	"fmt"
	"github.com/serveye/backend-tech/handlers/category"
	"github.com/serveye/backend-tech/pkg/config"
	"github.com/serveye/backend-tech/render"
	"log"
	"net/http"
)

const portNumber = ":7890"



func Excute() {
	http.HandleFunc("/admin", category.GetCategories)

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
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("server failed to start", err)
	}
}
