package models

import (
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}


// BeforeCreate will set a UUID rather than numeric ID.
//func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
//	b.ID = uuid.NewV4()
//	log.Println("Finally")
//	return
//}
