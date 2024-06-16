package types

import "github.com/golang-jwt/jwt/v5"

type JwtInterface struct {
	UserId string `json:"userId"`
	UserEmail string `json:"email"`
	UserStatus string `json:"userStatus"`
	jwt.MapClaims
}
