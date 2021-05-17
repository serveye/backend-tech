package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/models"
	"github.com/serveye/backend-tech/viewModels"
	"log"
	"net/http"
)

var app *config.AppConfig

// AppConfig sets the application wide config
func AppConfig(a *config.AppConfig) {
	app = a
}
func (m *Repository) GetCategories(w http.ResponseWriter, r *http.Request) {
	log.Println("came to controller")
	var categories []models.Category
	log.Println(config.DB)
	config.DB.Find(&categories)

	err := json.NewEncoder(w).Encode(categories)

	if err != nil {
		log.Println("Error occured while json parsing the categories")
	}
}

func (m *Repository) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c viewModels.CategoryCreate
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category := models.Category{Name: c.Name, Id: uuid.NewV4()}
	config.DB.Create(&category)
	log.Println(config.DB)
	err = json.NewEncoder(w).Encode(category)

	if err != nil {
		log.Println("Error occured while json parsing the categories")
	}
}
