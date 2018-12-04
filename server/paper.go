package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/leoryu/leo-ryu.herokuapp.com/model"
	"github.com/leoryu/leo-ryu.herokuapp.com/store"
)

func PubilushPaper(c echo.Context) error {
	paper := new(model.Paper)
	if err := c.Bind(paper); err != nil {
		return err
	}
	if c.Request().Method == echo.PUT {
		return store.ModifyPaper(c, paper)
	}
	return store.SavePaper(c, paper)
}

func DeletePaper(c echo.Context) error {
	id := c.Param("id")
	return store.DeletePaper(c, id)
}

func GetPaper(c echo.Context) error {
	id := c.Param("id")
	paper, err := store.GetPaper(c, id)
	if err != nil {
		return err
	}
	if paper == nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, paper)
}
