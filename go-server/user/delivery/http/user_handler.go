package http

import (
	"go-server/domain"

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
		return
	}
	c.JSON(200, gin.H{"data": users})
}

func (u *UserHandler) GetUserByUserID(c *gin.Context) {
	userID := c.Param("userID")
	logrus.Debug("userID:", userID)

	anUser, err := u.UserUsecase.GetByID(c, userID)
	if err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Query digimon error",
		// })
		return
	}

	c.JSON(200, swagger.User{
		Id:     anUser.Id,
		Name:   anUser.Name,
		Email:  anUser.Email,
		Rating: anUser.Rating,
	})
}
