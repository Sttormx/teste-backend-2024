package main

import (
	"ms-go/app/messager"
	"ms-go/app/services/consumer"
	_ "ms-go/db"
	"ms-go/router"
	"time"
)

func main() {
	time.Sleep(time.Second * 10)

	// Mensageria
	instance := messager.New()
	go instance.SetupReader(messager.TOPIC_RAILS_TO_GO, consumer.Handler)

	router.Run()
}
