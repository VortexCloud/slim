package services

import (
	"context"
	"log"
	"slim/app/libs"
)

type TestService struct{}

func (s TestService) GetValue() string {
	redis := libs.GetInstance(0)
	val, err := redis.Get(context.Background(), "golang").Result()
	if err != nil {
		log.Println("redis获取数据失败")
		panic(err)
	}
	log.Println("redis获取数据成功")
	return val
}

func (s TestService) GetUserList() interface{} {
	//user := &models.User{}

	return ""
}
