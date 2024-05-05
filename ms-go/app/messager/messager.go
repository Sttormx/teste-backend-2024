package messager

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	TOPIC_RAILS_TO_GO = "rails-to-go"
	TOPIC_GO_TO_RAILS = "go-to-rails"
	PARTITION_DEFAULT = 0
)

type Messager struct {
	Conn *kafka.Conn
}

func New() *Messager {
	return &Messager{}
}

func (c *Messager) Setup(topic string, partition int) *Messager {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:29092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	c.Conn = conn
	return c
}

func (c *Messager) Close() {
	if err := c.Conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func (c *Messager) Write(message []byte) (*Messager, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return c, err
	}

	_, err = c.Conn.Write(message)
	return c, err
}

func (c *Messager) SetupReader(topic string, consumer func(message []byte)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:29092"},
		Topic:     topic,
		Partition: PARTITION_DEFAULT,
	})

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			break
		}

		consumer(msg.Value)
	}
}
