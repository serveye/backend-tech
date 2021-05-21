package User

import (
	"github.com/serveye/backend-tech/models"
)

type BaseUser struct {
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Email       string         `json:"email`
	PhoneNumber string         `json:"phone_number"`
	Address     models.Address `json:"address"`
}
