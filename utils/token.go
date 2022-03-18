package utils

import (
	"test_sql/config"
	"time"

	"test_sql/app/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type UserClaims struct {
	UserId      int    `json:"userId"`
	Role        string `json:"role"`
	AccessUuid  string `json:"accessUuid"`
	RefreshUuid string `json:"refreshUuid"`
	jwt.StandardClaims
}

func CreateToken(userId int) (*model.TokenDetails, error) {
	var err error
	modelTokenDetail := &model.TokenDetails{}

	id, _ := uuid.NewV4()

	// access token will expire after 15 minute
	modelTokenDetail.AtExpires = time.Now().Add(time.Duration(config.Get().JwtToken.AccessTokenMaxAge) * time.Second).Unix()
	modelTokenDetail.AccessUuid = id.String()
	// refresh token will expire after 7 days
	modelTokenDetail.RtExpires = time.Now().Add(time.Duration(config.Get().JwtToken.RefreshTokenMaxAge) * time.Second).Unix()
	modelTokenDetail.RefreshUuid = id.String()
	accessKey := []byte(config.Get().JwtToken.AccessTokenSecretKey)
	refreshKey := []byte(config.Get().JwtToken.RefreshTokenSecretKey)

	// Create the Claims
	accessTokenClaims := UserClaims{
		UserId:     userId,
		AccessUuid: modelTokenDetail.AccessUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: modelTokenDetail.AtExpires,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	modelTokenDetail.AccessToken, err = token.SignedString(accessKey)

	refreshTokenClaims := UserClaims{
		UserId:      userId,
		RefreshUuid: modelTokenDetail.RefreshUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: modelTokenDetail.RtExpires,
			Issuer:    "test",
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	modelTokenDetail.RefreshToken, err = token.SignedString(refreshKey)

	if err != nil {
		return nil, err
	}
	return modelTokenDetail, nil
}
