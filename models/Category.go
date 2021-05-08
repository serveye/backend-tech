package models

import (
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Base
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

func (Category) TableName() string {
	return "category"
}
