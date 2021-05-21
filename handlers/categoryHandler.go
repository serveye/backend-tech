package handlers

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/helpers/paginations"
	"github.com/serveye/backend-tech/models"
	"github.com/serveye/backend-tech/viewModels"
	"log"
	"net/http"
	"strconv"
)

var app *config.AppConfig

// AppConfig sets the application wide config
func AppConfig(a *config.AppConfig) {
	app = a
}
func (m *Repository) GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	orderQuery := r.URL.Query().Get("orderBy")

	p := &paginations.Param{
		DB:      config.DB,
		Page:    page,
		Limit:   limit,
		OrderBy: orderQuery,
	}
	paginatedData, err := paginations.Pagging(p, &categories)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(paginatedData)

	if err != nil {
		log.Println("Error occured while json parsing the categories")
	}
	//config.DB.Find(&categories)
	//
	//err := json.NewEncoder(w).Encode(categories)
	//
	//if err != nil {
	//	log.Println("Error occured while json parsing the categories")
	//}
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
