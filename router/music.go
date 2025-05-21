package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	music_api "github.com/Dream-ming/myMusic/api/music"
)

func music_router(r *gin.Engine) {
	r.GET("/music", func(c *gin.Context) {
		c.HTML(http.StatusOK, "music.html", nil)
	})

	api := r.Group("/api") 
	{
		songs := api.Group("/songs") 
		{
			songs.GET("/history_top", music_api.GetHistoryTopSongs)
			songs.GET("/today_top", music_api.GetTodayTopSongs)
		}
 	}
}