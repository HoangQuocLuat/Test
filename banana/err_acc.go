package banana

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại")
	RegisFail = errors.New("Đăng ký thất bại")
)