package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{"message": "page not found"})
}
