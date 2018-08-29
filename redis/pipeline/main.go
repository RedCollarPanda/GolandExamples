package main

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	pingResult, err := redisClient.Ping().Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ping result = ", pingResult, err)




	redisClient.WrapProcessPipeline(func(old func([]redis.Cmder) error) func([]redis.Cmder) error {
		return func(cmds []redis.Cmder) error {
			fmt.Printf("pipeline starting processing: %v\n", cmds)
			err := old(cmds)
			fmt.Printf("pipeline finished processing: %v\n", cmds)
			return err
		}
	})

	redisClient.Pipelined(func(pipe redis.Pipeliner) error {
		pipe.Set("hello", "world", 0)
		pipe.Ping()
		pipe.Ping()
		return nil
	})



}
