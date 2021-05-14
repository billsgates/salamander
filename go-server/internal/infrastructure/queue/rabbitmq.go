package queue

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// RabbitMQHandler defines the RabbitMQ handler
type RabbitMQHandler struct {
	conn    *amqp.Connection
	queue   amqp.Queue
	channel *amqp.Channel
}

// NewRabbitMQHandler creates new RabbitMQHandler
func NewRabbitMQHandler() *RabbitMQHandler {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		logrus.Fatal(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		logrus.Fatal(err)
	}

	queue, err := channel.QueueDeclare(
		"verification", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            //arguments
	)
	if err != nil {
		logrus.Fatal(err)
	}

	return &RabbitMQHandler{
		conn:    conn,
		queue:   queue,
		channel: channel,
	}
}

// Conn returns the conn property
func (r RabbitMQHandler) Conn() *amqp.Connection {
	return r.conn
}

// Queue returns the queue property
func (r RabbitMQHandler) Queue() amqp.Queue {
	return r.queue
}

// Channel returns the channel property
func (r RabbitMQHandler) Channel() *amqp.Channel {
	return r.channel
}
