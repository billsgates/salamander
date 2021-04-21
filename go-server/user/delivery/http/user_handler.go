package http

import (
	"go-server/domain"
	"net/http"

	swagger "go-server/go"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(e *gin.RouterGroup, authMiddleware gin.HandlerFunc, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

	userEndpoints := e.Group("users", authMiddleware)
	{
		userEndpoints.GET("", handler.GetAllUsers)
		userEndpoints.GET("/:userID", handler.GetUserByUserID)
	}
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := u.UserUsecase.FetchAll(c)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *UserHandler) GetUserByUserID(c *gin.Context) {
	userID := c.Param("userID")

	anUser, err := u.UserUsecase.GetByID(c, userID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, swagger.User{
		Id:     anUser.Id,
		Name:   anUser.Name,
		Email:  anUser.Email,
		Rating: anUser.Rating,
	})
}
