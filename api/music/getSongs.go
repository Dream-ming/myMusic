package music

import (
    "net/http"
    "github.com/Dream-ming/myMusic/internal/music"
    "github.com/gin-gonic/gin"
)

func GetHistoryTopSongs(c *gin.Context) {
    songs, err := music.GetHistoryTopSongs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, songs)
}

func GetTodayTopSongs(c *gin.Context) {
    songs, err := music.GetTodayTopSongs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, songs)
}
