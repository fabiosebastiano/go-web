package main

import (
	"github.com/fabiosebastiano/go-web/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	routers.Register(e)
	e.Run(":8080")
}
