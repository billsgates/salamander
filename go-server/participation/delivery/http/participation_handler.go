package http

import (
	"go-server/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ParticipationHandler struct {
	RoomUsecase domain.RoomUsecase
}

func NewParticipationHandler(e *gin.RouterGroup, authMiddleware gin.HandlerFunc, roomUsecase domain.RoomUsecase) {
	handler := &ParticipationHandler{
		RoomUsecase: roomUsecase,
	}

	roomEndpoints := e.Group("participant", authMiddleware)
	{
		roomEndpoints.DELETE("", handler.LeaveRoom)
	}
}

func (p *ParticipationHandler) LeaveRoom(c *gin.Context) {
	var body domain.ParticipationRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := p.RoomUsecase.LeaveRoom(c, body.RoomId, body.UserId)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
