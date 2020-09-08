package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goprovision/api/handlers"
)

func main() {
	//setup gin framework
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/v1/working", handlers.IsApplicationWorking) //api route

	r.Run(":80") //run application
}
