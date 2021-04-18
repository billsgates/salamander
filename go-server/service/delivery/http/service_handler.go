package http

import (
	"go-server/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServiceHandler struct {
	serviceUsecase domain.ServiceUsecase
}

func NewServiceHandler(e *gin.RouterGroup, serviceUsecase domain.ServiceUsecase) {
	handler := &ServiceHandler{
		serviceUsecase: serviceUsecase,
	}

	userEndpoints := e.Group("services")
	{
		userEndpoints.GET("", handler.GetAllServices)
	}
}

func (s *ServiceHandler) GetAllServices(c *gin.Context) {
	services, err := s.serviceUsecase.FetchAll(c)
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(200, gin.H{"data": services})
}
