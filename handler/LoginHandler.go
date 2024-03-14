package handler

// import (
// 	"net/http"
// 	"thuchanh_go/controllers"
// 	"thuchanh_go/types"

// 	"github.com/gin-gonic/gin"
// )

// func LoginUserHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		//lay data input
// 		var req types.GetUserReq
// 		var err = ctx.ShouldBindJSON(&req)
// 		if err != nil {
// 			panic(err)
// 		}

// 		//xử lý dữ liệu và đưa ra kết quả
// 		res, err := controllers.GetUserLogic(req)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, nil)
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, res)
// 	}
// }