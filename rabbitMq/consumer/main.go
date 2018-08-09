package main

import (
	"github.com/streadway/amqp"
	"gopkg.in/gcfg.v1"
	"log"
)

type Config struct {
	Section_rabbit struct {
		RabbitUrl  string
		RabbitQeue string
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func handleMessages(bundle string) string {

	return "{ok}"
}

func run_consumer() {

	//conn, err := amqp.Dial(cfg.Section_rabbit.RabbitUrl)
	//failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := globalConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.Section_rabbit.RabbitQeue, // name
		true, // durable
		//false,       // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			log.Println("got message: " + string(d.Body))
			result := handleMessages(string(d.Body))

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          []byte(result),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever

}

func initAmqp(rabbitUrl string) {
	var err error
	globalConn, err = amqp.Dial(cfg.Section_rabbit.RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
}

var (
	cfg        Config
	globalConn *amqp.Connection
)

func main() {

	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Section_rabbit.RabbitUrl != "" {

		initAmqp(cfg.Section_rabbit.RabbitUrl)
		defer globalConn.Close()
	} else {

		log.Println("No rabbitmq configuration found.")
		return
	}

	forever := make(chan bool)

	//fixme
	for i := 0; i < 3; i++ {
		go run_consumer()
	}

	<-forever
}
