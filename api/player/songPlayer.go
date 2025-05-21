package api

import (
	"net/http"
	"strconv"

	"github.com/Dream-ming/myMusic/internal/music"
	"github.com/gin-gonic/gin"
)

func PlayerAPIHandler(c *gin.Context) {
	// 获取歌曲 ID
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song ID"})
		return
	}

	// 转换 ID 为 uint64
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	// 获取歌曲详细信息
	song, err := music.GetSongByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// 如果歌曲未找到，返回 404
	if song == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	// 返回歌曲信息
	c.JSON(http.StatusOK, song)
}