package music

import (
    "database/sql"
	"errors"

    "github.com/Dream-ming/myMusic/initialize"
)

// var DB = initialize.DB

type Song struct {
	ID          uint64         `json:"id"`
	Name        string         `json:"name"`        // 可为 NULL
	Duration    sql.NullInt64  `json:"duration"`    // 可为 NULL
	OssURL      string         `json:"oss_url"`     // 假设不允许 NULL
	LyricPath   sql.NullString `json:"lyric_path"`  // 可为 NULL
	ReleaseYear sql.NullString `json:"release_year"`// 可为 NULL
	Artist      string         `json:"artist"`      // 可为 NULL
	PlayCount   uint64         `json:"play_count"`  // 假设默认值为 0
}

func GetHistoryTopSongs() ([]Song, error) {
    rows, err := initialize.DB.Query(`
        SELECT s.id, s.name, s.artist, s.oss_url, spc.count
        FROM song s
        JOIN song_play_counter spc ON s.id = spc.song_id
        ORDER BY spc.count DESC LIMIT 10
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var songs []Song
    for rows.Next() {
        var s Song
        if err := rows.Scan(&s.ID, &s.Name, &s.Artist, &s.OssURL, &s.PlayCount); err != nil {
            return nil, err
        }
        songs = append(songs, s)
    }
    return songs, nil
}

func GetTodayTopSongs() ([]Song, error) {
    rows, err := initialize.DB.Query(`
        SELECT s.id, s.name, s.artist, s.oss_url, spc.count
        FROM song s
        JOIN song_play_counter spc ON s.id = spc.song_id
        ORDER BY spc.count DESC LIMIT 10
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var songs []Song
    for rows.Next() {
        var s Song
        if err := rows.Scan(&s.ID, &s.Name, &s.Artist, &s.OssURL, &s.PlayCount); err != nil {
            return nil, err
        }
        songs = append(songs, s)
    }
    return songs, nil
}

func GetSongByID(id uint64) (*Song, error) {
	// 查询 song 表和 song_play_counter 表的数据
	row := initialize.DB.QueryRow(`
		SELECT 
			s.id, s.name, s.duration, s.oss_url, s.lyric_path, s.release_year, s.artist, 
			IFNULL(spc.count, 0) AS play_count
		FROM 
			song AS s
		LEFT JOIN 
			song_play_counter AS spc 
		ON 
			s.id = spc.song_id
		WHERE 
			s.id = ?
	`, id)

	var song Song
	err := row.Scan(&song.ID, &song.Name, &song.Duration, &song.OssURL, &song.LyricPath, &song.ReleaseYear, &song.Artist, &song.PlayCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err // 其他错误
	}

	// 增加播放计数
	err = incrementPlayCount(id)
	if err != nil {
		return nil, err // 返回更新播放计数的错误
	}

	// 返回歌曲信息
	return &song, nil
}

func incrementPlayCount(songID uint64) error {
	// 使用 INSERT ON DUPLICATE KEY UPDATE 实现自增
	_, err := initialize.DB.Exec(`
		INSERT INTO song_play_counter (song_id, count) 
		VALUES (?, 1) 
		ON DUPLICATE KEY UPDATE count = count + 1
	`, songID)

	return err
}
