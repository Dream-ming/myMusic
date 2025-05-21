package user

import(
    "database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
    "github.com/Dream-ming/myMusic/initialize"
)

func Login(username, password string) (uint64, error) {
	var id uint64
	var hashedPwd string
	err := initialize.DB.QueryRow("SELECT id, password_hash FROM user WHERE username=?", username).Scan(&id, &hashedPwd)
	if err == sql.ErrNoRows {
		return 0, errors.New("用户名不存在")
	} else if err != nil {
		return 0, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) != nil {
		return 0, errors.New("密码错误")
	}
	return id, nil
}