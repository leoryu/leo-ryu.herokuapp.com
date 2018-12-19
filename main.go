package main

import "github.com/leoryu/leo-ryu.herokuapp.com/router"

func main() {
	app := router.Router()
	log.Fatal(app.Run(":7070"))
}

