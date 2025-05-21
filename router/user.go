package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func user_router(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/api/login", func(c *gin.Context) {
		var req struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "参数错误"})
            return
        }

        // 假设用户名=admin，密码=123456 作为示例
        if req.Username == "admin" && req.Password == "123456" {
            // 登录成功
            c.JSON(http.StatusOK, gin.H{"success": true})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "用户名或密码错误"})
        }
	})
}