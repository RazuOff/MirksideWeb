package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",
		gin.H{})
}
