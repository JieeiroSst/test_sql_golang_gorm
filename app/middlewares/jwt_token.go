package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"test_sql/app/model"
	ormredis "test_sql/app/orm/redis"
	"test_sql/config"
	"test_sql/utils"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserGinKey = "CurrentUserId"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.ParseWithClaims(
		tokenString,
		&utils.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().JwtToken.AccessTokenSecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*model.TokenDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*utils.UserClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%d", claims.UserId), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.TokenDetails{
			AccessUuid: claims.AccessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth, err := ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		fmt.Printf("tokenAuth %+v \n ", tokenAuth)
		userId, err := ormredis.Token.Get(tokenAuth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		c.Set(UserGinKey, userId)
		c.Next()
	}
}