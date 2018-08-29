package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/gcfg.v1"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Section_basic struct {
		Type string
		IP   string
		Port string
	}

	Section_rabbit struct {
		RabbitUrl  string
		RabbitQeue string
	}
}

type limitHandler struct {
	connc   chan struct{}
	handler http.Handler
}

func (h *limitHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	select {
	case <-h.connc:

		fmt.Println("ServeHTTP")
		h.handler.ServeHTTP(w, req)

		h.connc <- struct{}{}
	default:
		http.Error(w, "503 too busy", http.StatusServiceUnavailable)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func NewLimitHandler(maxConns int, handler http.Handler) http.Handler {
	h := &limitHandler{
		connc:   make(chan struct{}, maxConns),
		handler: handler,
	}
	for i := 0; i < maxConns; i++ {
		h.connc <- struct{}{}
	}
	return h
}

func newTestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("newTestHandler")
		handler(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("handler ")

	if b, err := ioutil.ReadAll(r.Body); err == nil {

		fmt.Println("message is ", string(b))

		res := publishMessages(string(b))

		w.Write([]byte(res))
	} else {
		log.Println("Error while message reading.")
	}
}

func publishMessages(payload string) string {

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := strconv.FormatInt(time.Now().UnixNano(), 10)

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // noWait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	defer ch.QueueDelete(queueName, false, false, true)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Println("Consume ok")
	corrId := randomString(32)

	log.Println("corrId ", corrId)

	err = ch.Publish(
		"", // exchange
		cfg.Section_rabbit.RabbitQeue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode:  amqp.Transient,
			ContentType:   "application/json",
			CorrelationId: corrId,
			Body:          []byte(payload),
			Timestamp:     time.Now(),
			ReplyTo:       q.Name,
		})

	failOnError(err, "Failed to Publish on RabbitMQ")
	log.Println("Publish ok")

	result := "ERROR"
	for d := range msgs {
		if corrId == d.CorrelationId {
			failOnError(err, "Failed to convert body to integer")
			log.Println("result = ", string(d.Body))
			result = string(d.Body)
			return result
		} else {
			log.Println("waiting for result = ")
		}
	}

	log.Println("return")

	return result
}

func initAmqp(rabbitUrl string) {
	var err error
	conn, err = amqp.Dial(cfg.Section_rabbit.RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
}

func panicMiddlewareHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

var (
	cfg  Config
	conn *amqp.Connection
)

func PrintConfig(cfg Config) {
	fmt.Println("got config")
	fmt.Println("Type = ", cfg.Section_basic.Type)
	fmt.Println("IP = ", cfg.Section_basic.IP)
	fmt.Println("Port = ", cfg.Section_basic.Port)
	fmt.Println("RabbitUrl = ", cfg.Section_rabbit.RabbitUrl)
	fmt.Println("RabbitQeue = ", cfg.Section_rabbit.RabbitQeue)
}

func main() {

	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	PrintConfig(cfg)

	if cfg.Section_rabbit.RabbitUrl != "" {

		initAmqp(cfg.Section_rabbit.RabbitUrl)
		defer conn.Close()
	} else {

		log.Println("No rabbitmq configuration found.")
		return
	}

	mux := http.NewServeMux()

	handlr := NewLimitHandler(10099, newTestHandler())
	panicWrapper := panicMiddlewareHttp(handlr)

	mux.Handle("/messages/rabbit", panicWrapper)
	server := http.Server{
		Addr:         cfg.Section_basic.Port,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println(server.ListenAndServe())
}
