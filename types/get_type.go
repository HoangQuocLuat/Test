package types

type GetUserReq struct {
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
}



type GetUserRes struct {
	ID           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
	Email        string `json:"email"`
	Result       Result `json:"result"`
	Token		 Token 	`json:"token"`
}

type Result struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Token  string `json:"token"`
}

type PutUserReq struct {
	Hashpassword string `json:"hashpassword"`
	Email        string `json:"email"`
}

type GetInfor struct {
	ID           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
}