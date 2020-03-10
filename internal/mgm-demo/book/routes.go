package book

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	book := route.Group("/book")
	book.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "spam",
		})
	})
}
