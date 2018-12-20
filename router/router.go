package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leoryu/leo-ryu.herokuapp.com/config"
	"github.com/leoryu/leo-ryu.herokuapp.com/router/middleware"
	"github.com/leoryu/leo-ryu.herokuapp.com/server"
	"github.com/leoryu/leo-ryu.herokuapp.com/store/datastore"
)

func Router() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	api := e.Group("/api")
	api.Use(middleware.Store(datastore.New()))
	api.POST("/login", server.Login)
	api.GET("/paper/:id", server.GetPaper)
	api.GET("/introductions", server.GetIntroductions)
	{
		admin := api.Group("/admin")
		admin.Use(middleware.Auth(config.GetSecret()))
		admin.POST("/paper", server.CreatePaper)
		admin.PUT("/paper/:id", server.EditPaper)
		admin.DELETE("/paper/:id", server.DeletePaper)
	}
	return e
}

