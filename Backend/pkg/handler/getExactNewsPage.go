package handler

import (
	"mirkside/pkg/model"
	"mirkside/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type htmlForm struct {
	Header       string
	ContentBlock []model.ExactNewsBlock
}

func GetExactNewsPage(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	exactNews, err := repository.GetExactNewsPageById(idInt)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}
	shortNews, err := repository.GetShortNewsById(idInt)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	form := htmlForm{Header: shortNews.ContentHeader, ContentBlock: exactNews}

	c.HTML(http.StatusOK, "exactPage.html", form)
}
