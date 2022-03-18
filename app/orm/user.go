package orm

import (
	"test_sql/app/model"

	"github.com/jinzhu/gorm"
)

type userOrm struct {
	db *gorm.DB
}

type IUser interface {
	Create(user *model.User) (err error)
	GetByPhoneNumber(phoneNumber string) (user *model.User, err error)
	GetById(userId int) (user *model.User, err error)
}

var User IUser

func init() {
	User = &userOrm{}
}

// this function help to mock different database. example: for testing
func InitUserOrm(db *gorm.DB) *userOrm {
	return &userOrm{db}
}

func (o *userOrm) Create(user *model.User) (err error) {
	result := o.db.Create(user)
	return result.Error
}

func (o *userOrm) GetByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User
	result := o.db.Unscoped().Model(&model.User{}).
		Where("phone_number = ?", phoneNumber).
		Order("id DESC").
		Limit(1).
		Find(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, result.Error
}

func (o *userOrm) GetById(userId int) (*model.User, error) {
	var user model.User
	result := o.db.Unscoped().Model(&model.User{}).
		Where("id = ?", userId).
		Limit(1).
		Find(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, result.Error
}
