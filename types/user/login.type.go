package types

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
