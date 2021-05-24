package scheduler

import (
	"context"
	"time"

	"go-server/domain"
	adapterqueue "go-server/internal/adapter/queue"
	helper "go-server/internal/common"
	"go-server/internal/infrastructure/queue"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	cron           *cron.Cron
	roomUsecase    domain.RoomUsecase
	contextTimeout time.Duration
	producer       adapterqueue.Producer
}

func NewScheduler(roomUsecase domain.RoomUsecase, timeout time.Duration, queue *queue.RabbitMQHandler) *Scheduler {
	taipei, _ := time.LoadLocation("Asia/Taipei")
	c := cron.New(cron.WithLocation(taipei))
	s := &Scheduler{
		cron:           c,
		roomUsecase:    roomUsecase,
		contextTimeout: timeout,
		producer:       adapterqueue.NewProducer(queue.Channel(), "paymentCheck"),
	}
	// every 9:30 each day
	s.cron.AddJob("30 9 * * ?", s)
	logrus.Info(s.cron.Entries())
	logrus.Info("Start cron")
	c.Start()
	return s
}

func (s *Scheduler) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	// add function to cron
	s.sendEmailToStartingMembers(ctx)
	s.sendEmailToPaymentDueMembers(ctx)
}

// private, find room member that the starting date is today
func (s *Scheduler) sendEmailToStartingMembers(c context.Context) (err error) {
	logrus.Info("sendEmailToStartingMembers")
	members, err := s.roomUsecase.GetTodayStartingMember(c)
	if err != nil {
		return err
	}
	logrus.Info(members)
	for _, member := range members {
		message := helper.EncodeToBytes(&member)
		message = helper.Compress(message)
		s.producer.Publish(message)
	}
	return nil
}

// private, find room member that the payment date is due
func (s *Scheduler) sendEmailToPaymentDueMembers(c context.Context) (err error) {
	logrus.Info("sendEmailToPaymentDueMembers")
	members, err := s.roomUsecase.GetTodayPaymentDueMember(c)
	if err != nil {
		return err
	}
	logrus.Info(members)
	for _, member := range members {
		message := helper.EncodeToBytes(&member)
		message = helper.Compress(message)
		s.producer.Publish(message)
	}
	return nil
}
