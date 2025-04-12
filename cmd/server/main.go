package main

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	"github.com/Dream-ming/myMusic/initialize"
)

func main() {

	// r := gin.Default()
	// r.GET("/ping",func(c *gin.Context){
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"code":200,
	// 		"msg":"pong",
	// 	})
	// })
	// r.Run(":8888")

	initialize.InitAll()

}