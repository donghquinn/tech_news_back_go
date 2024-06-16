package types

type MiddlewareErrorResponse struct {
	Result bool `json:"result"`
	Code string `json:"code"`
	Message string `json:"message"`
}
