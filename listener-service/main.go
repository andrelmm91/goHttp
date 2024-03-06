package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	
	// start listening
	log.Println("listening for and consuming rabbitmq messages...")

	// create consumer 
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}
	
	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	// onlz continue when rabbitmq is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq") //rabbitmq should be the same as in docker file
		if err != nil {
			fmt.Println("RabbitMQ not yet ready")
			counts++
		} else {
			log.Println("connected to RabbitMQ")
			connection = c
			break
		}

	if counts > 5 {
		fmt.Println(err)
		return nil, err
	}

	backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
	log.Println("backing off...")
	time.Sleep(backoff)
	continue
	} 

	return connection, nil
}