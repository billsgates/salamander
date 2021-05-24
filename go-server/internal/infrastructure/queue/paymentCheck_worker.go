package queue

import (
	"go-server/domain"
	"go-server/internal/infrastructure/mail"

	helper "go-server/internal/common"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type PaymentCheckWorker struct {
	GmailHandler    *mail.GmailHandler
	RabbitMQHandler *RabbitMQHandler
	msgs            <-chan amqp.Delivery
}

func NewPaymentCheckWorker(rabbitMQHandler *RabbitMQHandler, gmailHandler *mail.GmailHandler) *PaymentCheckWorker {
	err := rabbitMQHandler.Channel().Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := rabbitMQHandler.Channel().Consume(
		"paymentCheck", // queue
		"",             // consumer
		false,          // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	return &PaymentCheckWorker{
		RabbitMQHandler: rabbitMQHandler,
		GmailHandler:    gmailHandler,
		msgs:            msgs,
	}
}

func (w *PaymentCheckWorker) Start() {
	forever := make(chan bool)

	go func() {
		for d := range w.msgs {
			message := helper.Decompress(d.Body)
			user := helper.DecodeToParticipationInfo(message)
			w.SendEmail(&user)
			d.Ack(false)
		}
	}()

	<-forever
}

func (w *PaymentCheckWorker) SendEmail(info *domain.ParticipationInfo) {
	data := struct {
		ReceiverName    string
		ServiceProvider string
		Plan            string
		AdminName       string
		AdminEmail      string
	}{
		ReceiverName:    info.UserName,
		ServiceProvider: info.ServiceProvider,
		Plan:            info.Plan_name,
		AdminName:       info.AdminName,
		AdminEmail:      info.AdminEmail,
	}
	status, err := w.GmailHandler.SendEmailOAUTH2("frankchen93011@gmail.com", data, "payment_due_template.txt")
	if err != nil {
		logrus.Info(err)
	}
	if status {
		logrus.Info("Email sent successfully using OAUTH")
	}
}
