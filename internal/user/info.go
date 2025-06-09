package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Dream-ming/myMusic/initialize"
)

// UserInfo 是返回给前端的结构体
type UserInfo struct {
	UserID         int    `json:"user_id"`
	Username       string `json:"user_name"`
	RegisteredAt   string `json:"registered_at"`
	DaysRegistered int    `json:"days_registered"`
}

// GetUserInfo 通过 userID 获取用户信息
func GetUserInfo(userID int) (*UserInfo, error) {
	var username string
	var createdAt time.Time

	err := initialize.DB.QueryRow("SELECT username, created_at FROM user WHERE id = ?", userID).
		Scan(&username, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	days := int(time.Since(createdAt).Hours() / 24)

	return &UserInfo{
		UserID:         userID,
		Username:       username,
		RegisteredAt:   createdAt.Format("2025-05-10 15:04:05"),
		DaysRegistered: days,
	}, nil
}