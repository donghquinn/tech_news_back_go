package types

type SignupRequest struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	Uuid string `json:"uuid"`
}

type AccountItem struct {
	Uuid string `json:"uuid"`
}