package model

import "gopkg.in/dgrijalva/jwt-go.v3"

type User struct {
	Username string `json:"username"`
	Is_Login bool   `json:"is_login"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
