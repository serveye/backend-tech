package User

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City string `json:"city"`
	Province string `json:"province"`
	PostalCode string `json:"postal_code"`
}

type User struct {
	firstName string
	lastName string
	email string
	phoneNumber string
    Address Address
	IsAuthenticated bool `json:"is_authenticated"`
	IsProvider bool `json:"is_provider"`

}
