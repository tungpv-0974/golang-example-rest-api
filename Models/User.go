package Models

import "example.com/m/v2/Config"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (b *User) TableName() string {
	return "users"
}

func CreateAUser(user *User) (error error) {
	if error = Config.DB.Create(user).Error; error != nil {
		return error
	}
	return nil
}

func GetAUser(user *User, id string) (error error) {
	if error := Config.DB.Where("id = ?", id).First(user).Error; error != nil {
		return error
	}
	return nil
}

func FindByUserName(user *User, userName string) (error error) {
	if error := Config.DB.Where("user_name = ?", userName).First(user).Error; error != nil {
		return error
	}
	return nil
}
