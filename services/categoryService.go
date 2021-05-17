package services

import (
	"github.com/serveye/backend-tech/config"
	"github.com/serveye/backend-tech/models"
	"github.com/serveye/backend-tech/viewModels"
)

func getCategories() []models.Category {
	var categories []models.Category
	config.DB.Find(&categories)
	return categories
}

func createCategory(categoryToCreate *viewModels.CategoryCreate) *viewModels.CategoryCreate {
	config.DB.Create(&categoryToCreate)
	return categoryToCreate
}
