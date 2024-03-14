package logic

// import (
// 	"log"
// 	"thuchanh_go/database"
// 	"thuchanh_go/models"
// 	"thuchanh_go/types"
// 	"thuchanh_go/utils"

// 	"github.com/zeromicro/go-zero/core/logx"
// 	"golang.org/x/crypto/bcrypt"
// )

// const SecretKey = "secret"

// func GetUserLogic(req types.GetUserReq) (*types.GetUserRes, error) {
// 	//connect db mysql
// 	db, err := database.DBConn()
// 	defer db.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	logx.Info(req)

// 	//get data
// 	var user models.UserTbl

// 	err = db.QueryRow("SELECT id, fullname, hashpassword, email  FROM user_tbl WHERE name = ?", req.Name).
// 		Scan(&user.ID, &user.Fullname, &user.Hashpassword, &user.Email)

// 	logx.Info(user)
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Hashpassword), []byte(req.Hashpassword))
// 	if err != nil {
// 		logx.Error(err)
// 		return &types.GetUserRes{
// 			Result: types.Result{
// 				Code:    400,
// 				Message: "sai mật khẩu",
// 			},
// 		}, nil
// 	}

// 	token, err := utils.GenerateJWT(user.ID, SecretKey)
// 	logx.Info(err)
// 	if err != nil {
// 		return &types.GetUserRes{
// 			Result: types.Result{
// 				Code:    400,
// 				Message: "không thể đăng nhập",
// 			},
// 		}, nil
// 	}

// 	return &types.GetUserRes{
// 		ID:       user.ID,
// 		Fullname: user.Fullname,
// 		Email:    user.Email,
// 		Result:   types.Result{Code: 200, Message: "thành công"},
// 		Token:    types.Token{Token: token},
// 	}, nil
// }
