package main

import (
	"github.com/go-redis/redis"
	"log"
	"reflect"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pingResult, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ping result = ", pingResult)

	key := "hello"

	//set
	err = redisClient.Set(key, "world", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	intResult, err := redisClient.Exists(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(intResult)
	if intResult == 0 {
		log.Fatal("No value for key: ", key)
	}

	result, err := redisClient.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	//mset
	msetResulr, err := redisClient.MSet("key1", "val1", "key2", "val2" ).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(msetResulr)
	keyResult := redisClient.Keys("key*").Val()
	log.Println(keyResult)

	if !reflect.DeepEqual(keyResult,[]string{"key1", "key2"}) {
		log.Println(keyResult)
		log.Println([]string{"key1", "key2"})
		log.Fatal("Deepequal missmatch")
	}
}
