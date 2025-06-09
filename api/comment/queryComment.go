package comment

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Dream-ming/myMusic/internal/comment"
)

func QueryComment(c *gin.Context) {
	var req struct {
		SongID  int    `json:"song_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效", "details": err.Error()})
		return
	}

	// 调用内部逻辑获取评论
	comments, err := comment.QueryComment(req.SongID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询评论失败", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}
