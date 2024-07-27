package types

type ResponseDto struct {
	ResCode int `json:"resCode"`
	DataRes ResponseResult `json:"dataRes"`
}

type ResponseResult struct {
	Result any `json:"result"`
}