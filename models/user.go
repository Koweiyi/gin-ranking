package models

import (
	"go-ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.DB.Create(&user).Error
	return user.Id, err
}

func GetUserInfo(userId int) (User, error) {
	var user User
	err := dao.DB.Where("id = ?", userId).First(&user).Error
	return user, err
}
