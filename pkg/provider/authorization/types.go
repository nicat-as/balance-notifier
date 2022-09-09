package authorization

import "time"

type KapitalAuthorization struct {
}
type Token struct {
	RefreshToken string
	ExpireTime   time.Time
}

type AuthResponse struct {
	RefreshToken string `json:"jwtrefreshtoken"`
	AccessToken  string `json:"jwttoken"`
}

type RefreshTokenRequest struct {
	Token string `json:"refreshToken"`
}

const bearer = "Bearer "

const Username = "USERNAME"
const Password = "PASSWORD"
