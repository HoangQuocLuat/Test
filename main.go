package main

import (
	"thuchanh_go/database"
	handler "thuchanh_go/handler/account"
	logic "thuchanh_go/logic/account"
	router "thuchanh_go/router/acc"

	"github.com/gin-gonic/gin"
)

func main() {
	sql := &database.Sql{
		UserName: "postgres",
		Password: "HL123",
		DbName:   "test",
	}
	sql.Connect()
	defer sql.Close()

	// init server
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
