package jwt

import (

)


type Config struct {
	Secret string `json:"secret" validate:"required"`
	AccsessPeriod int `json:"access_period" validate:"required"`
	RefreshPeriod int `json:"refresh_period" validate:"required"`
}

// jwt.json
// {
// 	"secret": "0de81fe3867deeejghn6369124ca1077",
// 	"access_period": 60,
// 	"refresh_period": 120
// }

type UserData struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type TokenData struct {
	User *UserData `json:"user"`
	Exp int `json:"exp"`
}


type TokensSet struct{
	Access string `json:"access"`
	Refresh string `json:"refresh"`
}


type JWT struct {
	config *Config
}

func New (config *Config) *JWT {

	jwt := JWT{config: config}

	return  &jwt
}

func (jwt *JWT) CreateTokens(user *UserData) (*TokensSet, error) {
	
	// accessTokenData := TokenData{
	// 	User: user,
	// 	Exp: int(time.Now().Unix()) + jwt.config.AccsessPeriod,
	// }

	tokensSet := TokensSet{}

	return &tokensSet, nil
}

func (jwt *JWT) UpdateTokens(refreshToken string) (*TokensSet, error) {
	tokensSet := TokensSet{}

	return  &tokensSet, nil
}

