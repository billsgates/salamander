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

func NewAuthHandler(e *gin.Engine, authUsecase domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: authUsecase,
	}

	e.POST("/api/v1/auth/signup", handler.SignUp)
	e.POST("/api/v1/auth/signin", handler.SignIn)
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var body swagger.SignupRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
		return
	}

	anUser, err := h.AuthUsecase.SignUp(c.Request.Context(), body.Name, body.Email, body.Password)
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(201, swagger.User{
		Id:             anUser.Id,
		Name:           anUser.Name,
		Email:          anUser.Email,
		PasswordDigest: anUser.PasswordDigest,
		Rating:         anUser.Rating,
		CreatedAt:      anUser.CreatedAt,
		UpdatedAt:      anUser.UpdatedAt,
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var body swagger.LoginRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		// c.JSON(500, &swagger.ModelError{
		// 	Code:    3000,
		// 	Message: "Internal error. Parsing failed",
		// })
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

	c.JSON(200, swagger.LoginResponse{Token: token})
}
