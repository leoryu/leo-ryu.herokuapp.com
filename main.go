package main

import "github.com/leoryu/leo-ryu.herokuapp.com/router"

func main() {
	app := router.Router()
	app.Run(":7070")
}

