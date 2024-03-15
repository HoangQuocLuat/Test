package main

import (
	"thuchanh_go/database"
	handler "thuchanh_go/handler/account"
	logic "thuchanh_go/logic/account"
	"thuchanh_go/redis"
	router "thuchanh_go/router/acc"

	"github.com/gin-gonic/gin"
)

func main() {
	//init sql
	sql := &database.Sql{
		UserName: "postgres",
		Password: "HL123",
		DbName:   "test",
	}
	sql.Connect()
	defer sql.Close()
	//init redis
	redis := &redis.Redis{
		Addr: 		"localhost:6379",
		Password:   "",   
	}
	redis.Connect()
	
	r := gin.New()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Type"},
	// 	AllowCredentials: true,
	// }))

	// router

	// r.POST("/user/register", handler.RegisHandler())
	// run server

	userHandler := handler.AccountHandler{
		UserLogic: logic.NewAccRegisterLogic(sql),
	}
	api := router.API{
		Gin:        r,
		AccHandler: userHandler,
	}
	api.SetupRoute()

	r.Run("0.0.0.0:8888")
}
