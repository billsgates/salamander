package http

import (
	"go-server/domain"

	swagger "go-server/go"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoomHandler struct {
	RoomUsecase domain.RoomUsecase
}

func NewRoomHandler(e *gin.RouterGroup, authMiddleware gin.HandlerFunc, RoomUsecase domain.RoomUsecase) {
	handler := &RoomHandler{
		RoomUsecase: RoomUsecase,
	}

	roomEndpoints := e.Group("rooms", authMiddleware)
	{
		roomEndpoints.POST("", handler.CreateRoom)
	}
}

func (u *RoomHandler) CreateRoom(c *gin.Context) {
	var body swagger.RoomCreateRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
		return
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)

	room, err := u.RoomUsecase.Create(c.Request.Context(), &domain.Room{
		MaxCount:     body.MaxCount,
		AdminId:      user.Id,
		ServiceId:    body.ServiceId,
		PlanName:     body.PlanName,
		StartingTime: time.Now(),
		EndingTime:   time.Now(),
	})
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(200, swagger.Room{
		Id:              room.Id,
		AccountName:     room.AccountName,
		AccountPassword: room.AccountPassword,
		StartingTime:    room.StartingTime,
		EndingTime:      room.EndingTime,
		CreatedAt:       room.CreatedAt,
		UpdatedAt:       room.UpdatedAt,
		MaxCount:        room.MaxCount,
		AdminId:         room.AdminId,
		ServiceId:       room.ServiceId,
		PlanName:        room.PlanName,
	})
}
