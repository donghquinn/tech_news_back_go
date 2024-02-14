package types

type SignupRequest struct {
	Email string `json:"email" xml:"email" binding:"min=5,required"`
	Name string `json:"name" xml:"name" binding:"min=2,required"`
	Password string `json:"password" xml:"password" binding:"min=5,max=15,required"`
}

type LoginRequest struct {
	Email string `json:"email" xml:"email" binding:"min=5,required"`
	Password string `json:"password" xml:"password" binding:"min=5,max=15,required"`
}

type LogoutRequest struct {
	Uuid string `json:"uuid" xml:"uuid" binding:"min=8,required"`
}

type AccountItem struct {
	Uuid string `json:"uuid" binding:"required"`
}