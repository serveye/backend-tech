package User

type CreateUser struct {
	BaseUser BaseUser `json:"base_user"`
	Password string   `json:"password"`
}
