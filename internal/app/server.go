package server

import (
	"github.com/Koubae/httpMaker/internal/app/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Start() {
	var appName = os.Getenv("APP_NAME")
	if appName == "" {
		appName = config.AppNameDefault
	}
	// ------------------------ < Error Handler > -------------------- \\
	defer func() {
		log.Printf("%s Shutting Down...", appName)
		if revived := recover(); revived != nil {
			log.Fatalf("%s Panicked: %+v\n", appName, revived)
		}
	}()
	// ------------------------ < !!!!!!!!!!!!!! > -------------------- \\

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = config.AppPortDefault
	}

	app := gin.New()
	app.LoadHTMLGlob(config.PathIndex)
	app.Static(config.PathAssetsRelative, config.PathAssets)

	app.Use(config.LoggerWithFormatter())
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Websiste",
		})
	})

	log.Printf("/!| App  %s Running /!| Listening and Serving at 0.0.0.0:%s", appName, port)
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func Configure() {
	log.Println("Go Initializing App ... ")

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading the .env file!")
	}

	// Starting gin app
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
		if ginMode == "debug" {
			log.Println("Setting Gin App to Debug!!!")
		}
	}
	gin.ForceConsoleColor()
}
