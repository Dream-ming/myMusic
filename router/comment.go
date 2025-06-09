package router

import (
	"github.com/gin-gonic/gin"

	comment_api "github.com/Dream-ming/myMusic/api/comment"
	token_api "github.com/Dream-ming/myMusic/api/token"
)

func comment_router(r *gin.Engine) {
	api := r.Group("/api") 
	{
		songs := api.Group("/comment") 
		{
			songs.POST("/query_comment", comment_api.QueryComment)
		}
	}
	
	commentGroup := r.Group("/api/comment")
	commentGroup.Use(token_api.JWTAuthMiddleware())
	{
		commentGroup.POST("/push_comment", comment_api.PushComment)
	}
}