package types

type ResponseDto struct {
	ResCode int `json:"resCode" xml:"resCode" binding:"required"`
	DataRes ResponseResult `json:"dataRes" xml:"dataRes" binding:"required"`
}

type ResponseResult struct {
	Result any `json:"result" xml:"result" binding:"required"`
}