package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/Ramshackle-Jamathon/go-quickPerm"
	goredis "github.com/go-redis/redis"
	"github.com/kataras/go-sessions"
	"github.com/kataras/go-sessions/sessiondb/redis"
	"github.com/kataras/go-sessions/sessiondb/redis/service"
	"github.com/prometheus/common/log"
	"github.com/twinj/uuid"
	"gopkg.in/gcfg.v1"
	"io/ioutil"
	"strconv"
	"strings"
)

type DataArray struct {
	Set []int `json:"set"`
}

type Config struct {
	Basic struct {
		MaxPayloadSize int
		ListeningPort  string
	}
}

//var maxPayloadSize = 16777216

func main() {

	var cfg Config
	//Лучше передавать путь до конфига, но для простоты можно и так
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	//В конфиг выше можно так же вынести все эти параметры
	db := redis.New(service.Config{Network: service.DefaultRedisNetwork,
		Addr:        service.DefaultRedisAddr,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      "",
	})

	//так как менеджер сессий использует redis, то мы можем также его использовать для
	//очереди перестановок
	redisClient := goredis.NewClient(&goredis.Options{
		Addr:     service.DefaultRedisAddr,
		Password: "",
		DB:       0,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sess := sessions.New(sessions.Config{Cookie: "sessionscookieid"})
	sess.UseDatabase(db)

	app := http.NewServeMux()
	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(
			"You should navigate to the `/v1/init` or `/v1/next` instead\n")))
	})

	//Handler /init/
	app.HandleFunc("/v1/init", func(w http.ResponseWriter, r *http.Request) {

		//при условии, что "..элементы множества могут быть только
		//положительные целые числа.." мы не делаем проверку, что это числа

		strLength, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Content length error!\n"))
			return
		}

		/*
			...размер множества ограничен 16Мб данных в теле запроса...

		*/
		if strLength > cfg.Basic.MaxPayloadSize {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Content length too big!\n"))
			return
		}

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal error!\n"))
			return
		}

		var data DataArray
		if err := json.Unmarshal(buf, &data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Json handling Internal error!\n"))
			return
		}

		s := sess.Start(w, r)

		uuid := uuid.NewV4().String()

		first := ""
		for permutation := range quickPerm.GeneratePermutationsInt(data.Set) {
			fmt.Println(permutation)

			answerData := DataArray{Set:permutation}
			result, err := json.Marshal(answerData)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Json handling Internal error!\n"))
				return
			}

			if first == "" {
				first = string(result)
				//...В ответе передается первая перестановка...
				continue
			}
			redisClient.RPush(uuid, string(result))
		}

		s.Set("uuid", uuid)
		w.Write([]byte(fmt.Sprintf("%s\n", first)))
	})

	//handler /next
	app.HandleFunc("/v1/next", func(w http.ResponseWriter, r *http.Request) {

		name := sess.Start(w, r).GetString("uuid")
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("uuid not found!\n"))
			return
		}

		w.Write([]byte(fmt.Sprintf("The name on the /set was: %s \n", name)))

		s := strings.Split(redisClient.RPop(name).String(), " ")
		if s[2] == "redis:" {
			var empty []int
			w.Write([]byte(fmt.Sprintf("%d\n", empty)))
			return
		}
		w.Write([]byte(fmt.Sprintf("%s\n", s[2])))
	})

	http.ListenAndServe(cfg.Basic.ListeningPort, app)
}

//curl -v -X POST 127.0.0.1:8080/init -H 'content-type: application/json' -d '{ "set": [1,2,3] }'

//curl -v -X GET 127.0.0.1:8080/next -H 'Cookie: sessionscookieid=*sessionid*'
