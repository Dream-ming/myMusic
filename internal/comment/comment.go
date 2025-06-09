package comment

import (
	"time"
)

// Comment 评论模型
type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	SongID    int       `json:"song_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}