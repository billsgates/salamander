package http

import (
	"net/http"

	"go-server/auth"
	"go-server/domain"
	swagger "go-server/go"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(e *gin.RouterGroup, authUsecase domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: authUsecase,
	}

	authEndpoints := e.Group("auth")
	{
		authEndpoints.POST("/signup", handler.SignUp)
		authEndpoints.POST("/signin", handler.SignIn)
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var body swagger.SignupRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.AuthUsecase.SignUp(c.Request.Context(), body.Name, body.Email, body.Password)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusCreated)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var body swagger.LoginRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.AuthUsecase.SignIn(c.Request.Context(), body.Email, body.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, swagger.LoginResponse{Token: token})
}
