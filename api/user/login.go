package user

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/Dream-ming/myMusic/internal/user"
	"github.com/Dream-ming/myMusic/internal/jwt"
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
	token, err := jwt.GenerateToken(uint64(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "生成JWT错误"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg": "登录成功", 
		"user_id": userID,
		"user_name": req.Username,
		"token": token,
	})
}