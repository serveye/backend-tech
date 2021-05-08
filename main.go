package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/controllers"
	"github.com/serveye/backend-tech/models"
	"github.com/spf13/viper"
	"net/http"
)
const portNumber = ":8080"

func main() {

	// viper config

	// Set the filename of the configuration file
	viper.SetConfigName("config")

	// Set the path to look for the configuration file
	viper.AddConfigPath(".")

	// enable viper to read env variables
	viper.AutomaticEnv()

	// set config file type
	viper.SetConfigType("yml")

	var appConfig config.AppConfig
	if err:=viper.ReadInConfig(); err!=nil {
		fmt.Printf("Error reading config file %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "dev_db")
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

config.SetEnvConfig(&appConfig)
	config.ConnectDataBase()
	router := mux.NewRouter()
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/check", func (w http.ResponseWriter, r *http.Request) {
		_, err :=fmt.Fprint(w, "Check Page")
		if err != nil {
			fmt.Println(err)
		}
	}).Methods("GET")

	_ = config.DB.AutoMigrate(models.Category{})
	err = http.ListenAndServe(portNumber, router)
	if err != nil {
		fmt.Println("server failed to start", err)
	}
}
