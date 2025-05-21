package user

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/Dream-ming/myMusic/internal/user"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	userID, err := user.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// 示例：返回用户ID，可以改为JWT token等
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg": "登录成功", 
		"user_id": userID,
	})
}