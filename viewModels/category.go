package viewModels

import "time"

type baseCategory struct {
	Name string `json:"name"`
}

type CategoryCreate struct {
	baseCategory
}

type CategoryView struct {
	baseCategory
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool `json:"is_deleted"`
}