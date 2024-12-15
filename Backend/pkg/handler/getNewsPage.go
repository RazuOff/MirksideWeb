package handler

import (
	"mirkside/pkg/model"
	"mirkside/pkg/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsBlock struct {
	NewsContent model.ShortNews
	Url         string
}

func GetNewsPage(c *gin.Context) {
	news := repository.GetShortNews()
	var newsBlocks []NewsBlock
	for _, item := range news {
		newsBlocks = append(newsBlocks, NewsBlock{NewsContent: item, Url: "/news/" + strconv.Itoa(item.ID)})
	}

	c.HTML(http.StatusOK, "newsPage.html", newsBlocks)
}
