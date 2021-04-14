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
	e.GET("/api/v1/users/:userID", handler.GetUserByUserID)
	e.PUT("/api/v1/users/:userID", handler.UpdateUserByUserID)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var body swagger.User
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
		return
	}

	anUser := domain.User{
		Name:  body.Name,
		Email: body.Email,
	}

	if err := u.UserUsecase.Create(c, &anUser); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Store failed",
		// })
		return
	}

	c.JSON(200, nil)
}

func (u *UserHandler) GetUserByUserID(c *gin.Context) {
	userID := c.Param("userID")

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
		Id:     anUser.ID,
		Name:   anUser.Name,
		Email:  anUser.Email,
		Rating: anUser.Rating,
	})
}

func (u *UserHandler) UpdateUserByUserID(c *gin.Context) {
	userID := c.Param("userID")

	anUser, err := u.UserUsecase.GetByID(c, userID)
	if err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Query digimon error",
		// })
		return
	}

	var body swagger.User
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
		return
	}

	anUser.Name = body.Name
	anUser.Email = body.Email

	if err := u.UserUsecase.Update(c, &anUser); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Store failed",
		// })
		return
	}
	c.JSON(200, nil)
}
