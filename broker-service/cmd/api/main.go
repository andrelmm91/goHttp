package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()


	app := Config{
		Rabbit: rabbitConn
	}

	log.Printf("Starting broker service on port %s \n", webPort)

	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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