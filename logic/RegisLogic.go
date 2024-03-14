package logic

// import (
// 	"log"
// 	"thuchanh_go/database"
// 	"thuchanh_go/types"

// 	"github.com/zeromicro/go-zero/core/logx"
// 	"golang.org/x/crypto/bcrypt"
// )

// func RegisLogic(req types.UserRegisReq) (*types.Result, error) {
// 	db, err := database.DBConn()
// 	defer db.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	logx.Info(req)

// 	// var count int
// 	// err = db.QueryRow("SELECT COUNT(*) FROM user_tbl WHERE name = ?", req.Username).Scan(&count)

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// if count > 0 {
// 	// 	// Tài khoản đã tồn tại, trả về lỗi
// 	// 	return &types.Result{
// 	// 		Code:    400,
// 	// 		Message: "tài khoản đã tồn tại",
// 	// 	}, nil
// 	// }

// 	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

// 	query := "INSERT INTO user_tbl (id, name, phone, email, username, password) VALUES(?, ?, ?, ?, ?, ?)"
// 	_, err = db.Exec(query, req.ID, req.Name, req.Phone, req.Email, req.Username, hashpassword)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &types.Result{
// 		Code:    200,
// 		Message: "thành công",
// 	}, nil
// }
