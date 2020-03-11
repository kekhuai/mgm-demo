package book

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	book := route.Group("/book")
	book.GET("", func(c *gin.Context) {
		c.JSON(200, Book{})
	})
	book.POST("", func(c *gin.Context) {
		newBook := &Book{
			Name:  "test",
			Pages: 10,
		}
		_ = mgm.Coll(newBook).Create(newBook)
		c.JSON(201, newBook)
	})
}
