package User

type User struct {
	BaseUser BaseUser `json:"base_user"`
	IsAuthenticated bool `json:"is_authenticated"`
	IsProvider bool `json:"is_provider"`
}
