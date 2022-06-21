package user

import "gohub/pkg/database"

//IsEmailExist 判断Email已被注册
func IsEmailExist(email string) bool {
	var count int64

	database.DB.Model(User{}).Where("email = ?", email).Count(&count)

	return count > 0
}

//IsPhoneExist 判断Email已被注册
func IsPhoneExist(phone string) bool {
	var count int64

	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}
