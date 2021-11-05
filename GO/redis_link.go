package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	pong,err := client.Ping().Result()
	fmt.Println(pong,err)
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val,err2 := client.Get("key").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("key",val)

	_ = client.HMSet("2009",map[string]interface{}{
		"main_uid": 2009,
		"friend_uid": 2008,
		"note": "我是你爸爸",
		"state":0,
	}).Err()
	maps,err5 := client.HGetAll("2009").Result()
	if err5 != nil {
		panic(err5)
	}
	println(maps["main_uid"])

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2",val2)
	}
}