package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"forum/global"
	"forum/models"
)

const secret = "mande.versatile"

func CheckUserExit(user *models.User) (error error) {
	var count int64
	result := global.Db.Model(&models.User{}).Where("user_name =?", user.UserName).Count(&count)
	if result.Error != nil {
		panic(result.Error)
	}
	if count > 0 {
		return errors.New(ErrorUserExit)
	}
	return
}

func Login(user *models.User) (error error) {
	// 记录用户原始密码
	originPassword := user.Password
	// 根据用户名查询数据库用户 信息
	var newUser *models.User
	global.Db.Model(&models.User{}).Where("user_name =?", user.UserName).First(&newUser)
	// 密码比对
	password := EncryptPassword([]byte(originPassword))
	if newUser.Password != password {
		return errors.New(ErrorPasswordWrong)
	}
	return nil
}

func EncryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
