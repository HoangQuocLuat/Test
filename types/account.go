package types


type GetIdUserReq struct {
	UserID int64 `uri:"user_id"`
}

type UserLoginReq struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
}
type UserLoginRes struct {
	Name         string `json:"name"`
	Hashpassword string `json:"hashpassword"`
}