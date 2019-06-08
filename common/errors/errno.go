package errors

const (
	ErrCustom        int32 = 9999
	ErrInvalidToken  int32 = 1000
	ErrUserNotExists int32 = 1001
	ErrPasswordError int32 = 1002
)

var ErrMsg = map[int32]string{
	ErrInvalidToken:  "无效的令牌",
	ErrUserNotExists: "用户不存在",
	ErrPasswordError: "密码错误",
}
