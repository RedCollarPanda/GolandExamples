package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:"localhost:6379",
			Password:"",
			DB:0,
		})

		pingResult, err := redisClient.Ping().Result()
		log.Println("ping result = ", pingResult, err)

		//TODO



}
