package user

import(
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Dream-ming/myMusic/internal/user"
)

func GetUserInfo(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未获取到用户信息"})
		return
	}

	userID, ok := userIDVal.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户ID类型转换失败"})
		return
	}
	
	userInfo, err := user.GetUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInfo)
}