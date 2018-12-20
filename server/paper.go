package server

import (
	"fmt"
	"net/http"
	"strconv"
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
	editCount, err := store.ModifyPaper(c, paper, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if editCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"err": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Edited paper %s", id)})
}

func DeletePaper(c *gin.Context) {
	id := c.Param("id")
	delCount, err := store.DeletePaper(c, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if delCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"err": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted Paper %s", id)})
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

func GetIntroductions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	count, introductions, err := store.GetIntroductions(c, limit, page)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "introductions": introductions})
}

