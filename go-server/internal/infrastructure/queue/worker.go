package queue

import (
	"github.com/sirupsen/logrus"
)

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}

func sendEmail(target string) {
	// m := gomail.NewMessage()
	// m.SetHeader("From", "admin@billsgate.com")
	// m.SetHeader("To", "kevin.ct.yu@ntu.im")
	// m.SetHeader("Subject", "Test")
	// m.SetBody("text/html", "Hello! Greetings from BillsGate!")

	// // d := gomail.NewDialer("mail.billsgate.com", 465, "", "")
	// // d := gomail.NewDialer("mailserver", 587, "admin", "password")
	// d := gomail.Dialer{Host: "localhost", Port: 587}

	// if err := d.DialAndSend(m); err != nil {
	// 	panic(err)
	// }
}

func NewWorker(rabbitMQHandler *RabbitMQHandler) {
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			logrus.Printf("Received a message: %s", d.Body)
			// sendEmail(string(d.Body))
			logrus.Printf("Done")
			d.Ack(false)
		}
	}()

	<-forever
}
