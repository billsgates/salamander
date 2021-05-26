package http

import (
	"go-server/domain"
	"net/http"
	"strconv"

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

	userEndpoints := e.Group("user", authMiddleware)
	{
		userEndpoints.GET("", handler.GetUser)
		// userEndpoints.GET("", handler.GetAllUsers)
		userEndpoints.PATCH("", handler.UpdateUserInfo)
		// userEndpoints.GET("/:userID", handler.GetUserByUserID)
		userEndpoints.GET("/:userID/rating", handler.GetUserRating)
		userEndpoints.PATCH("/:userID/rating", handler.UpdateUserRating)
	}
}

func (u *UserHandler) GetUser(c *gin.Context) {
	user := c.Value(domain.CtxUserKey).(*domain.User)
	userId := strconv.Itoa(int(user.Id))

	res, err := u.UserUsecase.GetByID(c, userId)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, domain.UserInfo{
		Id:     anUser.Id,
		Name:   anUser.Name,
		Email:  anUser.Email,
		Rating: anUser.Rating,
	})
}

func (u *UserHandler) UpdateUserInfo(c *gin.Context) {
	var body domain.UserRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.Value(domain.CtxUserKey).(*domain.User)
	body.Id = user.Id
	err := u.UserUsecase.Update(c, &body)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusCreated)
}

func (u *UserHandler) GetUserRating(c *gin.Context) {
	userID := c.Param("userID")
	res, err := u.UserUsecase.GetUserRating(c, userID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) UpdateUserRating(c *gin.Context) {
	var body domain.RatingRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userID := c.Param("userID")
	logrus.Info("UpdateUserRating userId", userID, body)
	err := u.UserUsecase.UpdateRating(c, userID, body.Rating)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}
