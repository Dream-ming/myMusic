package comment

import (
	"github.com/Dream-ming/myMusic/initialize"
)

func QueryComment(songID int) ([]Comment, error) {
	rows, err := initialize.DB.Query("SELECT c.id, c.user_id, u.username, c.song_id, c.content, c.created_at FROM comment c LEFT JOIN user u ON c.user_id = u.id WHERE song_id = ? ORDER BY created_at DESC", songID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.UserName, &comment.SongID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}