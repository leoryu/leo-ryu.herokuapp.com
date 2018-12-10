package router

import (
    "github.com/gin-gonic/gin"
    "github.com/leoryu/leo-ryu.herokuapp.com/server"
)

func Router() *gin.Engine {
    e := gin.New()
    e.Use(gin.Recovery())
    api := e.Group("/api")
    api.POST("/login", server.Login)
    api.GET("/paper/:id", server.GetPaper)
    {
        admin := api.Group("/admin")
        admin.POST("/paper", server.PubilushPaper)
        admin.PUT("/paper", server.PubilushPaper)
        admin.DELETE("/paper/:id", server.DeletePaper)
    }
    return e
}
