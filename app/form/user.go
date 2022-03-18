package form 

type RegisterPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber" validate:"nonzero,maxLen=20"`
	Password    string `json:"password" validate:"maxLen=11"`
}

type LoginPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber" validate:"nonzero,maxLen=20"`
	Password    string `json:"password" validate:"maxLen=11"`
}	