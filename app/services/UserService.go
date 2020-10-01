package services

import (
	"slim/app/libs"
	"slim/app/models"
)

const DBIndex int = 0

type UserService struct{}

func (userService UserService) GetValue() string {
	redis := libs.GetRedisInstance(DBIndex)
	val, err := redis.Get("golang").Result()
	if err != nil {
		return err.Error()
	}
	return val
}

func (userService UserService) GetUserList() interface{} {
	user := &models.User{}
	data := user.GetUserList()
	return data
}

func (userService UserService) CreateUser(name string, email string, password string) bool {
	user := models.User{}
	rst := user.CreateUser(name, email, password)
	return rst
}
