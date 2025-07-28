package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



type Config struct {
	Secret        string `json:"secret" validate:"required"`
	AccsessPeriod int    `json:"access_period" validate:"required"`
	RefreshPeriod int    `json:"refresh_period" validate:"required"`
}

// jwt.json
// {
// 	"secret": "0de81fe3867deeejghn6369124ca1077",
// 	"access_period": 60,
// 	"refresh_period": 120
// }


type TokenData struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
	RoleName string `json:"role_name"`
	Exp int64 `json:"exp"`
}

type TokensSet struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type JWT struct {
	config *Config
}

func New(config *Config) *JWT {

	jwt := JWT{config: config}

	return &jwt
}

func (j *JWT) CreateTokens(userID uint) (*TokensSet, error) {

	accessClaims := TokenData{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID: userID,
		Exp: time.Now().Add(time.Duration(j.config.AccsessPeriod)*time.Second).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessTokenString, err := accessToken.SignedString([]byte(j.config.Secret))

	if err != nil {
		return nil, err
	}

	refreshClaims := TokenData{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID: userID,
		Exp: time.Now().Add(time.Duration(j.config.RefreshPeriod)*time.Second).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(j.config.Secret))

	if err != nil {
		return nil, err
	}

	tokensSet := TokensSet{
		Access: accessTokenString,
		Refresh: refreshTokenString,
	}

	return &tokensSet, nil
}

func (j *JWT) ParseToken(tokenString string) (*TokenData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (any, error) {
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	tokenData := token.Claims.(*TokenData)

	if tokenData.Exp < time.Now().Unix() {
		return nil, errors.New("expire token")
	}

	return tokenData, nil	
}



func (j *JWT) UpdateTokens(refreshToken string) (*TokensSet, error) {

	tokenData, err := j.ParseToken(refreshToken)

	if err != nil {
		return nil, err
	}

	tokensSet, err := j.CreateTokens(tokenData.UserID)

	if err != nil {
		return  nil, err
	}

	return tokensSet, nil
}


