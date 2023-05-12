package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/get_logger", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is a test for logger!",
		})
	})
	r.Run()
}
