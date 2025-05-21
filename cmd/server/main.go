package main

import (
	// "encoding/json"
	// "net/http"

	"github.com/Dream-ming/myMusic/initialize"
	"github.com/Dream-ming/myMusic/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	
	router.RegisterRouters(r)

	initialize.InitAll()

	r.Run(":8888")

}
