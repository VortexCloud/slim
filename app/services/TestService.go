package services

import (
	"slim/app/libs"
	"slim/app/models"
)

const DBIndex int = 0

type TestService struct{}

func (s TestService) GetValue() string {
	redis := libs.GetRedisInstance(DBIndex)
	val, err := redis.Get("golang").Result()
	if err != nil {
		return err.Error()
	}
	return val
}

func (s TestService) GetUserList() interface{} {
	user := &models.User{}
	data := user.GetUserList()
	return data
}
