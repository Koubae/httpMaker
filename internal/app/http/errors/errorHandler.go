package errors

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorHandler404(app *gin.Engine) {
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})
}

func ErrorHandler405(app *gin.Engine) {
	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})
}

func ErrorHandlerCatchAll(app *gin.Engine) {
	app.Use(errorHandler)
}

func errorHandler(c *gin.Context) {

	if len(c.Errors) <= 0 {
		c.Next()
		return
	}
	log.Printf("Total Errors -> %d", len(c.Errors))
	for _, err := range c.Errors { // Todo: check what to do in case of errors
		log.Printf("Error -> %+v\n", err)
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR_SERVER", "message": "500 Internal Server Error"})
}
