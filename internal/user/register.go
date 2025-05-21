package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/Dream-ming/myMusic/initialize"
)

func Register(username, password string) error {
	// 检查用户名是否存在
	var exists int
	err := initialize.DB.QueryRow("SELECT COUNT(*) FROM user WHERE username=?", username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("用户名已存在")
	}
	// 加密密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = initialize.DB.Exec("INSERT INTO user(username, password_hash) VALUES (?, ?)", username, hashedPwd)
	return err
}