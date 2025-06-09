package comment

import(
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/Dream-ming/myMusic/internal/comment"
)

func PushComment(c *gin.Context) {
	var req struct {
		SongID  int    `json:"song_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效", "details": err.Error()})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录"})
		return
	}
	newComment, err := comment.PushComment(userID.(int), req.SongID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提交评论失败", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newComment)
}