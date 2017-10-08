package main

import (
	"gopie/config"
	"gopie/views"

	"github.com/gin-gonic/gin"
)

func urlConfig(engine *gin.Engine) {
	engine.POST("push/", views.Push)
}

func main() {
	config.Parse()
	engine := gin.Default()

	urlConfig(engine)

	engine.Run(config.Address)
}
