package types

type ResponseDto struct {
	ResCode int `json:"resCode" xml:"resCode" binding:"required"`
	DataRes any `json:"dataRes" xml:"dataRes" binding:"required"`
}