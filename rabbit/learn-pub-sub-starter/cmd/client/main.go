package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	gl "github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
)

func main() {
	const connectionString = "amqp://guest:guest@localhost:5672/"

	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	log.Println("Connection successful")

	gl.ClientWelcome()

	connectionChan, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	pubsub.PublishJSON(connectionChan, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	log.Println("Shutting down")
	fmt.Println("Starting Peril client...")
}
