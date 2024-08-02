package main

import (
	"gin-gonic/pkg/hello"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": hello.GetMessage(),
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
