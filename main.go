package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leoryu/leo-ryu.herokuapp.com/config"

	"github.com/leoryu/leo-ryu.herokuapp.com/router"
)

func main() {
	app := router.Router()
	gin.SetMode(config.Mode)
	log.Fatal(app.Run(":7070"))
}
