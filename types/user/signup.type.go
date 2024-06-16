package types

type SignupResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
}

type SignupErrorResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type SignupRequestStruct struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
