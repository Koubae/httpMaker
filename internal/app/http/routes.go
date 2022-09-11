package http

import (
	"github.com/Koubae/httpMaker/internal/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {

	/// -----------------
	// 	Controllers
	/// -----------------
	controllers.Index(app)

	/// -----------------
	// 	API-Resources
	/// -----------------
}
