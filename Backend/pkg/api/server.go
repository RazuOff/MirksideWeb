package api

import (
	"mirkside/pkg/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.Static("/css", "../../Frontend/css")
	router.Static("/img", "../../Frontend/img")
	router.Static("/js", "../../Frontend/js")
	router.LoadHTMLGlob("../../Frontend/html/*")

	fillEndpoints(router)

	router.Run()
}

func fillEndpoints(router *gin.Engine) {
	router.GET("/", handler.GetMainPage)
	router.GET("/news", handler.GetNewsPage)
	router.GET("/news/:id", handler.GetExactNewsPage)

	router.NoRoute(handler.GetNotFoundPage)
}
