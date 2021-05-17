package main

import (
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/handlers"
	"github.com/serveye/backend-tech/models"
	"github.com/serveye/backend-tech/router"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var appConfig config.AppConfig
	appConfig.IsProduction = false

	// Repos Creation
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	handlers.AppConfig(&appConfig)
	config.SetEnvConfig(&appConfig)
	config.ConnectDataBase()
	_ = config.DB.AutoMigrate(models.Category{})
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router.Routes(&appConfig),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
