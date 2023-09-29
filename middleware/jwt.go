package middleware

import "github.com/golang-jwt/jwt/v5"

type jwtCustomClains struct {
	UserId int    `json:"userId"`
	Name   string `json:"bool"`
	jwt.RegisteredClaims
}

func GenerateTokenJwt(UserId int, Name string) string {
	var claims = jwtCustomClains{
		UserId: UserId,
		Name:   Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte("secret"))
	return resultJWT
}
