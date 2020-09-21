package Models

import (
	"example.com/m/v2/Config"
)

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Role     Role
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
	if error := Config.DB.Preload("Role").Where("id = ?", id).First(user).Error; error != nil {
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

// fetch all users at once
func GetAllUser(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
