package User

import User2 "github.com/serveye/backend-tech/models/User"

type BaseUser struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email`
	PhoneNumber string `json:"phone_number"`
	Address User2.Address `json:"address"`
}
