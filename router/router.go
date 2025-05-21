package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func RegisterRouters(r *gin.Engine) {

	r.Use(favicon.New("./web/static/favicon/favicon.ico"))
	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "./web/static")

	index_router(r)
	user_router(r)
	music_router(r)
	player_router(r)
}