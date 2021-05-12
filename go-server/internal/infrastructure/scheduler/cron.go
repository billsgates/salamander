package scheduler

import (
	"context"
	"time"

	"go-server/domain"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	cron           *cron.Cron
	roomUsecase    domain.RoomUsecase
	contextTimeout time.Duration
}

func NewScheduler(roomUsecase domain.RoomUsecase, timeout time.Duration) *Scheduler {
	taipei, _ := time.LoadLocation("Asia/Taipei")
	c := cron.New(cron.WithLocation(taipei))
	s := &Scheduler{
		cron:           c,
		roomUsecase:    roomUsecase,
		contextTimeout: timeout,
	}
	// every 9:30 each day
	s.cron.AddJob("30 9 * * ?", s)
	logrus.Info(s.cron.Entries())
	logrus.Info("Start cron")
	c.Start()
	return s
}

// private, find room member that the starting date is today
func (s *Scheduler) sendEmailToStartingMembers(c context.Context) (err error) {
	logrus.Info("sendEmailToStartingMembers")
	members, err := s.roomUsecase.GetTodayStartingMember(c)
	if err != nil {
		return err
	}
	logrus.Info(members)
	return nil
}

func (s *Scheduler) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), s.contextTimeout)
	defer cancel()
	// add function to cron
	s.sendEmailToStartingMembers(ctx)
}
