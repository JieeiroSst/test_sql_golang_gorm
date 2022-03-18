package redis 

import (
	"github.com/go-redis/redis/v7"
	"test_sql/app/model"
	"strconv"
	"time"
)

type jwtToken struct {
	client *redis.Client
}

type IJwtToken interface {
	Create(userId uint, td *model.TokenDetails) (err error)
	Get(td *model.TokenDetails) (uint64, error)
}

var Token IJwtToken

func (o *jwtToken) Create(userId uint, td *model.TokenDetails) (err error) {
	accessTokenExpireAt := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	refreshTokenExpireAt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := o.client.Set(td.AccessUuid, strconv.Itoa(int(userId)), accessTokenExpireAt.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := o.client.Set(td.RefreshUuid, strconv.Itoa(int(userId)), refreshTokenExpireAt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (o *jwtToken) Get(tokenDetail *model.TokenDetails) (uint64, error) {
	userIdStr, err := o.client.Get(tokenDetail.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	return userId, nil
}