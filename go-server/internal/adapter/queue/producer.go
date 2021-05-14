package queue

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type producer struct {
	channel   *amqp.Channel
	queueName string
}

// NewProducer creates new producer with its dependencies
func NewProducer(ch *amqp.Channel, qn string) Producer {
	return producer{
		channel:   ch,
		queueName: qn,
	}
}

// Publish sends a Publishing from the client to an exchange on the server
func (p producer) Publish(message []byte) error {
	if err := p.channel.Publish(
		"",          // exchange
		p.queueName, // routing key
		false,       // mandatory
		false,
		amqp.Publishing{
			Headers:     amqp.Table{},
			ContentType: "text/plain",
			Body:        message,
		}); err != nil {
		logrus.Errorf("failed to publish message: %s", message)

		return err
	}

	logrus.Infof("new message publish: %s", message)

	return nil
}
