package services

import (
	"slim/app/libs"
	"slim/app/models"
)

type TestService struct{}

func (s TestService) GetValue() string {
	redis := libs.GetRedisInstance(1)
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
