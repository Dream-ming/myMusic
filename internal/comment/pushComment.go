package comment

import (
	"errors"
	"time"

	"github.com/Dream-ming/myMusic/initialize"
)

// PushComment 推送新的评论
func PushComment(userID, songID int, content string) (*Comment, error) {
	if content == "" {
		return nil, errors.New("评论内容不能为空")
	}

	// 插入评论
	result, err := initialize.DB.Exec("INSERT INTO comment (user_id, song_id, content, created_at) VALUES (?, ?, ?, ?)",
		userID, songID, content, time.Now())
	if err != nil {
		return nil, err
	}

	// 获取插入的评论 ID
	commentID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 返回新创建的评论
	return &Comment{
		ID:        int(commentID),
		UserID:    userID,
		SongID:    songID,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}