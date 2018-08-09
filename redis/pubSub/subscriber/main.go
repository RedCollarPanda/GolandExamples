package main

import (
	"github.com/go-redis/redis"
	"log"
	"reflect"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Section_redis struct {
		RedisUrl  string
		RedisQueue string
	}
}

func main() {

	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Fatal(err)
	}



	redisClient := redis.NewClient(&redis.Options{
		Addr:cfg.Section_redis.RedisUrl,
		Password:"",
		DB:0,
	})

	pingResult, err := redisClient.Ping().Result()
	log.Println("ping result: ", pingResult, err)

	pubsub := redisClient.Subscribe(cfg.Section_redis.RedisQueue)
	defer pubsub.Close()

	for {

		msg, err := pubsub.Receive()
		if err != nil {
			log.Println(err)
			break;
		}

		switch inner := msg.(type) {

		case *redis.Subscription:
			log.Println("Received subscription: ", inner.Channel)

		case *redis.Message:
			log.Println("Received message: ", inner.Payload)
		default:
			log.Println("default message: ", reflect.TypeOf(msg))


		}


	}




}
