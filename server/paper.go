package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
	"github.com/leoryu/leo-ryu.herokuapp.com/store"
)

func CreatePaper(c *gin.Context) {
	paper := new(model.Paper)
	if err := c.ShouldBindJSON(paper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paper.CreatedAt = time.Now().UTC()
	paper.EditedAt = time.Now().UTC()
	if err := store.SavePaper(c, paper); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
}

func EditPaper(c *gin.Context) {
	paper := new(model.Paper)
	id := c.Param("id")
	if err := c.ShouldBindJSON(paper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paper.EditedAt = time.Now().UTC()
	if err := store.ModifyPaper(c, paper, id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
}

func DeletePaper(c *gin.Context) {
	id := c.Param("id")
	if err := store.DeletePaper(c, id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
}

func GetPaper(c *gin.Context) {
	id := c.Param("id")
	paper, err := store.GetPaper(c, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if paper == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, paper)
}

