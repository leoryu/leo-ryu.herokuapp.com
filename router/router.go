package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
	"github.com/leoryu/leo-ryu.herokuapp.com/server"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	api := e.Group("/api")
	api.POST("/login", server.Login)
	api.GET("/paper/:id", server.GetPaper)
	{
		admin := api.Group("/admin", middleware.JWT(model.Secret))
		admin.POST("/paper", server.PubilushPaper)
		admin.PUT("/paper", server.PubilushPaper)
		admin.DELETE("/paper/:id", server.DeletePaper)
	}
	return e
}

