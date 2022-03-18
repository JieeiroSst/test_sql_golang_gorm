package response 


import (
	"test_sql/app/model"
)

type User struct {
	Id          uint   `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"createdAt"`
}

func PopulateUserView(user model.User) (User, error) {
	var userView User
	createdAt := user.CreatedAt.String()

	userView = User{
		Id:          user.ID,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   createdAt,
	}

	return userView, nil
}