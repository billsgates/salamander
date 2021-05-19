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
		roomEndpoints.GET("/:roomID", handler.GetRoomInfo)
		roomEndpoints.PATCH("/:roomID", handler.UpdateRoomInfo)
		roomEndpoints.DELETE("/:roomID", handler.DeleteRoom)
		roomEndpoints.POST("/:roomID/round", handler.AddRound)
		roomEndpoints.DELETE("/:roomID/round", handler.DeleteRound)
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

func (u *RoomHandler) DeleteRoom(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = u.RoomUsecase.Delete(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (u *RoomHandler) GetJoinedRooms(c *gin.Context) {
	rooms, err := u.RoomUsecase.GetJoinedRooms(c)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

func (u *RoomHandler) GetRoomInfo(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	roomInfo, err := u.RoomUsecase.GetRoomInfo(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotMember {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	roomRound, err := u.RoomUsecase.GetRound(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotMember {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	admin, err := u.RoomUsecase.GetRoomAdmin(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotMember {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	members, err := u.RoomUsecase.GetRoomMembers(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotMember {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	roomInfo.Admin = admin
	roomInfo.Round = roomRound
	roomInfo.Members = members

	c.JSON(http.StatusOK, roomInfo)
}

func (u *RoomHandler) GenerateInvitationCode(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	code, err := u.RoomUsecase.GenerateInvitationCode(c, int32(roomID))
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
		if err == room.ErrRoomFull {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
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

func (u *RoomHandler) UpdateRoomInfo(c *gin.Context) {
	var body domain.RoomRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	logrus.Info("update body: ", body)

	err = u.RoomUsecase.UpdateRoom(c, int32(roomID), &body)
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotHost {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			return
		}
		if err == room.ErrMaxCountExceed {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}

func (u *RoomHandler) DeleteRound(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = u.RoomUsecase.DeleteRound(c, int32(roomID))
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotHost {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			return
		}
		if err == room.ErrNoRound {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (u *RoomHandler) AddRound(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomID"), 10, 32)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var body domain.RoundRequest
	body.RoomId = int32(roomID)

	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = u.RoomUsecase.AddRound(c, int32(roomID), &body)
	if err != nil {
		logrus.Error(err)
		if err == room.ErrNotHost {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			return
		}
		if err == room.ErrRoundAlreadyCreated {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}
