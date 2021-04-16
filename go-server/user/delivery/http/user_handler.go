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

func NewUserHandler(e *gin.Engine, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

	e.POST("/api/v1/users", handler.CreateUser)
	e.GET("/api/v1/users", handler.GetAllUsers)
	e.GET("/api/v1/users/:userID", handler.GetUserByUserID)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var body swagger.UserRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
		return
	}

	aUser := domain.User{
		Name:  body.Name,
		Email: body.Email,
	}

	_, err := u.UserUsecase.Create(c, &aUser)
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(200, swagger.User{
		Id:        aUser.Id,
		Name:      aUser.Name,
		Email:     aUser.Email,
		Rating:    aUser.Rating,
		CreatedAt: aUser.CreatedAt,
		UpdatedAt: aUser.UpdatedAt,
	})
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
