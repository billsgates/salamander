/*
 * HermitCrab api server
 *
 * This is a sample server of HermitCrab app.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"fmt"
	_authHandlerHttpDelivery "go-server/auth/delivery/http"
	_authUsecase "go-server/auth/usecase"
	_invitationRepo "go-server/invitation/repository/mysql"
	_participationHandlerHttpDelivery "go-server/participation/delivery/http"
	_participationRepo "go-server/participation/repository/mysql"
	_roomHandlerHttpDelivery "go-server/room/delivery/http"
	_roomRepo "go-server/room/repository/mysql"
	_roomUsecase "go-server/room/usecase"
	_serviceHandlerHttpDelivery "go-server/service/delivery/http"
	_serviceRepo "go-server/service/repository/mysql"
	_serviceUsecase "go-server/service/usecase"
	_userHandlerHttpDelivery "go-server/user/delivery/http"
	_userRepo "go-server/user/repository/mysql"
	_userUsecase "go-server/user/usecase"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func sayPongJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	logrus.Info("HTTP server started")

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Taipei")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
			"Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowCredentials: true,
	}))
	r.GET("/", sayHello)
	r.GET("/ping", sayPongJSON)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _userRepo.NewmysqlUserRepository(db)
	roomRepo := _roomRepo.NewmysqlRoomRepository(db)
	serviceRepo := _serviceRepo.NewmysqlServiceRepository(db)
	participationRepo := _participationRepo.NewmysqlParticipationRepository(db)
	invitationRepo := _invitationRepo.NewmysqlInvitationRepository(db)

	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	roomUsecase := _roomUsecase.NewRoomUsecase(roomRepo, participationRepo, serviceRepo, invitationRepo, timeoutContext)
	serviceUsecase := _serviceUsecase.NewServiceUsecase(serviceRepo, timeoutContext)
	authUsecase := _authUsecase.NewAuthUseCase(
		userRepo,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"),
	)

	authMiddleware := _authHandlerHttpDelivery.NewAuthMiddleware(authUsecase)

	v1Router := r.Group("/api/v1/")
	{
		_authHandlerHttpDelivery.NewAuthHandler(v1Router, authUsecase)
		_userHandlerHttpDelivery.NewUserHandler(v1Router, authMiddleware, userUsecase)
		_roomHandlerHttpDelivery.NewRoomHandler(v1Router, authMiddleware, roomUsecase)
		_serviceHandlerHttpDelivery.NewServiceHandler(v1Router, serviceUsecase)
		_participationHandlerHttpDelivery.NewParticipationHandler(v1Router, authMiddleware, roomUsecase)
	}

	logrus.Fatal(r.Run(":" + viper.GetString("server.address")))
}
