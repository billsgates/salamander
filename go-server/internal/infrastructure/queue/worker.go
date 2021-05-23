package queue

import (
	"go-server/domain"
	"go-server/internal/infrastructure/mail"

	helper "go-server/internal/common"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}

type Worker struct {
	GmailHandler    *mail.GmailHandler
	RabbitMQHandler *RabbitMQHandler
	msgs            <-chan amqp.Delivery
}

func NewWorker(rabbitMQHandler *RabbitMQHandler, gmailHandler *mail.GmailHandler) *Worker {
	err := rabbitMQHandler.Channel().Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := rabbitMQHandler.Channel().Consume(
		"verification", // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	return &Worker{
		RabbitMQHandler: rabbitMQHandler,
		GmailHandler:    gmailHandler,
		msgs:            msgs,
	}
}

func (w *Worker) Start() {
	forever := make(chan bool)

	go func() {
		for d := range w.msgs {
			message := helper.Decompress(d.Body)
			user := helper.DecodeToUser(message)
			w.SendEmail(&user)
			d.Ack(false)
		}
	}()

	<-forever
}

func (w *Worker) SendEmail(user *domain.User) {
	data := struct {
		ReceiverName string
		SenderName   string
	}{
		ReceiverName: user.Name,
		SenderName:   "Bills Gate",
	}
	status, err := w.GmailHandler.SendEmailOAUTH2("kevinyu05062006@gmail.com", data, "sample_template.txt")
	if err != nil {
		logrus.Info(err)
	}
	if status {
		logrus.Info("Email sent successfully using OAUTH")
	}
}
