package libs

import "github.com/go-redis/redis/v8"

func GetInstance(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.20.1.3:6379",
		Password: "",
		DB:       db,
	})
	return client
}

//var ctx = context.Background()
//
//func ExampleNewClient()  {
//	rdb := redis.NewClient(&redis.Options{
//		Addr: "172.20.1.3:6379",
//		Password: "",
//		DB: 0,
//	})
//	pong,err := rdb.Ping(ctx).Result()
//	fmt.Println(pong,err)
//}
