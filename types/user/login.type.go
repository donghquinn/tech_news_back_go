package types

type LoginRequestStruct struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


type LoginResponseType struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Token string `json:"token"`
}

type LoginErrorResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type LoginRedisStruct struct {
	Name string `json:"name"`
	Email string `json:"email"`
	UserId string `json:"userId"`
}

type LoginUserQueryResult struct {
	Uuid string `json:"uuid"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	PasswordToken string `json:"passwordToken"`
	IsLogined bool `json:"isLogined"`
	SignedIn string `json:"signedIn"`
	Logined string `json:"logined"`
}