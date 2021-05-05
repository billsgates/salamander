package http

import (
	"go-server/domain"
	"go-server/room"
	"net/http"
	"strconv"

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
		roomEndpoints.GET("", handler.GetJoinedRooms)
		roomEndpoints.POST("/:roomID/invitation", handler.GenerateInvitationCode)
		roomEndpoints.POST("/join", handler.JoinRoom)
	}
}

func (u *RoomHandler) CreateRoom(c *gin.Context) {
	var body domain.RoomRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := u.RoomUsecase.Create(c, &body)
	if err != nil {
		logrus.Error(err)
		if err == room.ErrMaxCountExceed {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}

func (u *RoomHandler) GetJoinedRooms(c *gin.Context) {
	user := c.Value(domain.CtxUserKey).(*domain.User)

	rooms, err := u.RoomUsecase.GetJoinedRooms(c.Request.Context(), user.Id)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

func (u *RoomHandler) GenerateInvitationCode(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)

	code, err := u.RoomUsecase.GenerateInvitationCode(c.Request.Context(), int32(roomID), user.Id)
	if code == "" || err != nil {
		logrus.Error(err)
		if err == room.ErrNotHost {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func (u *RoomHandler) JoinRoom(c *gin.Context) {
	var body domain.RoomJoinRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := u.RoomUsecase.JoinRoom(c, body.InvitationCode)
	if err != nil {
		logrus.Error(err)
		if err == room.ErrInvalidInvitationCode {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		if err == room.ErrAlreadyJoined {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
