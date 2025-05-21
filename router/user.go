package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

    user_api "github.com/Dream-ming/myMusic/api/user"
)

func user_router(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

    api := r.Group("/api") 
	{
		user := api.Group("/user")
		{
			user.POST("/register", user_api.Register)
            user.POST("/login", user_api.Login)
			user.POST("")
		}	
 	}
}