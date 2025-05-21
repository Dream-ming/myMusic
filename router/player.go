package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

	player_api "github.com/Dream-ming/myMusic/api/player"
)

func player_router(r *gin.Engine) {
	r.GET("/player", func(c *gin.Context) {
		c.HTML(http.StatusOK, "player.html", nil)
	})
	api := r.Group("/api") 
	{
		player := api.Group("/player")
		{
			player.GET("", player_api.PlayerAPIHandler)
		}	
 	}
}