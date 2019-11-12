package main

import (
"blog/router"
"github.com/gin-gonic/gin"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	routers := router.InitRouter()
	routers.Run(":5489")
}
