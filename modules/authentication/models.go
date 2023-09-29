package authentication

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"admin"`
}

type Claim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtResponse struct {
	JwtToken string `json:"jwtToken"`
}
