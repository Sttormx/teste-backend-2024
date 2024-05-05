package main

import (
	"ms-go/app/messager"
	"ms-go/app/services/consumer"
	_ "ms-go/db"
	"ms-go/router"
	"time"
)

func main() {
	// Kafka is refusing connections when starting
	// To solve it, i put a delay
	time.Sleep(time.Second * 10)

	// Messaging between microsservices
	// Reader is in a new goroutine to not block API Router
	instance := messager.New()
	go instance.SetupReader(messager.TOPIC_RAILS_TO_GO, consumer.Handler)

	router.Run()
}
