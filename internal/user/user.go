package user

import (
	"github.com/Dream-ming/myMusic/internal/db/mysql"
)

func GetUserInfo(user ID int) (*model.User, error) {
	var u model.User
    row := initmysql.DB.QueryRow("SELECT id, username, age FROM users WHERE id = ?", userID)
    err := row.Scan(&u.ID, &u.Username, &u.Age)
    if err != nil {
        return nil, err
    }
    return &u, nil
}