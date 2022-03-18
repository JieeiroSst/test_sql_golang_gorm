package entity

import (
	"fmt"
	"test_sql/app/form"
	"test_sql/app/model"
	"test_sql/app/orm"
	"test_sql/app/orm/redis"
	"test_sql/utils"
)

type UserEntity struct{}

type IUserEntity interface {
	Create(input form.RegisterPhoneNumber) (*model.User, error)
}

func (e *UserEntity) Create(input form.RegisterPhoneNumber) (model.User, error) {
	var user model.User
	//TODO: normalize phone number

	existedUser, err := orm.User.GetByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return user, err
	}

	if existedUser != nil {
		return user, fmt.Errorf("phone number is existed")
	}

	password := utils.HashPassword(input.Password)
	if len(password) == 0 {
		return model.User{}, nil
	}

	user = model.User{
		PhoneNumber: input.PhoneNumber,
		Password:    password,
	}

	err = orm.User.Create(&user)

	return user, err
}

func (e *UserEntity) LoginPhone(input form.LoginPhoneNumber) (map[string]string, error) {
	var tokens map[string]string
	user, err := orm.User.GetByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return tokens, err
	}

	if user == nil {
		return tokens, fmt.Errorf("phone number is not exist")
	}

	// Validate password
	err = utils.ComparePassword(input.Password, user.Password)
	if err != nil {
		err = fmt.Errorf("username or password is not match")
		return tokens, err
	}

	redisToken, err := utils.CreateToken(int(user.ID))
	if err != nil {
		err = fmt.Errorf("can not generate token")
		return tokens, err
	}

	err = redis.Token.Create(user.ID, redisToken)
	if err != nil {
		return tokens, err
	}

	tokens = map[string]string{
		"access_token":  redisToken.AccessToken,
		"refresh_token": redisToken.RefreshToken,
	}

	return tokens, err
}

func (e *UserEntity) GetById(userId int) (model.User, error) {
	user, err := orm.User.GetById(userId)
	if err != nil {
		return model.User{}, err
	}
	if user == nil {
		return model.User{}, err
	}
	return *user, err
}
