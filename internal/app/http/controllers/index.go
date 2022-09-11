package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(app *gin.Engine) {

	index := app.Group("/")

	{
		index.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		index.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

}
