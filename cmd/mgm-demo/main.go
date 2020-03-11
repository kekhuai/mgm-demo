package main

import (
	"github.com/Kamva/mgm/v2"
	"github.com/gin-gonic/gin"
	"github.com/kekhuay/mgm-demo/internal/mgm-demo/book"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	_ = mgm.SetDefaultConfig(nil, "mgm-demo", options.Client().ApplyURI("mongodb://localhost:27017"))
	r := gin.Default()
	book.Routes(r)
	r.Run(":8080")
}
