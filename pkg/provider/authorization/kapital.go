package authorization

import (
	"github.com/nicat-as/balance-notifier/pkg/authorization"
	"github.com/nicat-as/balance-notifier/pkg/http"
	"github.com/nicat-as/balance-notifier/pkg/provider"
	"time"
)

var tokenMap = map[string]Token{}

func (k KapitalAuthorization) GetToken(req UsernamePassword) string {
	if req.JwtToken != nil {
		if v, ok := tokenMap[*req.JwtToken]; ok {
			if v.ExpireTime.Before(time.Now()) {
				return bearer + *req.JwtToken
			} else {
				token, refresh := Refresh(v.RefreshToken)
				delete(tokenMap, *req.JwtToken)
				tokenMap[token] = Token{
					RefreshToken: refresh,
					ExpireTime:   time.Now(),
				}
				return bearer + token
			}
		}
	}
	token, refresh := Auth(req)
	tokenMap[token] = Token{
		RefreshToken: refresh,
		ExpireTime:   time.Now(),
	}
	return bearer + token
}

func Auth(req UsernamePassword) (string, string) {
	res, err := http.DoRequest[UsernamePassword, provider.KapitalResponseData[AuthResponse]](
		provider.GetEndpoint(provider.KapitalLogin),
		http.Post, nil, req,
	)
	if err != nil {
		return "", ""
	}
	tokenData := res.ResponseData
	return tokenData.AccessToken, tokenData.RefreshToken
}

func Refresh(token string) (string, string) {
	res, err := http.DoRequest[RefreshTokenRequest, provider.KapitalResponseData[AuthResponse]](
		provider.GetEndpoint(provider.KapitalRefresh),
		http.Post, nil, RefreshTokenRequest{Token: token},
	)
	if err != nil {
		return "", ""
	}
	tokenData := res.ResponseData
	return tokenData.AccessToken, tokenData.RefreshToken
}

type UsernamePassword struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	JwtToken *string `json:"jwttoken"`
}

var _ authorization.Authorization[UsernamePassword, string] = (*KapitalAuthorization)(nil)
