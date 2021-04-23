package http

import (
	"go-server/domain"
	"net/http"

	swagger "go-server/go"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoomHandler struct {
	RoomUsecase domain.RoomUsecase
}

func NewRoomHandler(e *gin.RouterGroup, authMiddleware gin.HandlerFunc, roomUsecase domain.RoomUsecase) {
	handler := &RoomHandler{
		RoomUsecase: roomUsecase,
	}

	roomEndpoints := e.Group("rooms", authMiddleware)
	{
		roomEndpoints.POST("", handler.CreateRoom)
	}
}

func (u *RoomHandler) CreateRoom(c *gin.Context) {
	var body domain.RoomCreateRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)

	room, err := u.RoomUsecase.Create(c.Request.Context(), &domain.Room{
		MaxCount:  body.MaxCount,
		AdminId:   user.Id,
		ServiceId: body.ServiceId,
		PlanName:  body.PlanName,
	})
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, swagger.Room{
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
