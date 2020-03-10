package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kekhuay/mgm-demo/internal/mgm-demo/book"
)

func main() {
	r := gin.Default()
	book.Routes(r)
	r.Run(":8080")
}
