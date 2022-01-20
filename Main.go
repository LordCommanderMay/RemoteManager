package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/cpu_info", func(c *gin.Context) {
		c.JSON(200, cpuInfo())
	})
	r.GET("/cpu_usage", func(c *gin.Context) {
		c.JSON(200, cpuUsage())
	})
	r.GET("/ram_info", func(c *gin.Context) {
		c.JSON(200,  ramInfo())

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}