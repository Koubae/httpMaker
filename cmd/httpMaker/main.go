package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

//var db = make(map[string]string)

//func setupRouter() *gin.Engine {
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//
//	// Get user value
//	r.GET("/user/:name", func(c *gin.Context) {
//		user := c.Params.ByName("name")
//		value, ok := db[user]
//		if ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//		}
//	})
//
//	// Authorized group (uses gin.BasicAuth() middleware)
//	// Same than:
//	// authorized := r.Group("/")
//	// authorized.Use(gin.BasicAuth(gin.Credentials{
//	//	  "foo":  "bar",
//	//	  "manu": "123",
//	//}))
//	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//		"foo":  "bar", // user:foo password:bar
//		"manu": "123", // user:manu password:123
//	}))
//
//	/* example curl for /admin with basicauth header
//	   Zm9vOmJhcg== is base64("foo:bar")
//
//		curl -X POST \
//	  	http://localhost:8080/admin \
//	  	-H 'authorization: Basic Zm9vOmJhcg==' \
//	  	-H 'content-type: application/json' \
//	  	-d '{"value":"bar"}'
//	*/
//	authorized.POST("admin", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		// Parse JSON
//		var json struct {
//			Value string `json:"value" binding:"required"`
//		}
//
//		if c.Bind(&json) == nil {
//			db[user] = json.Value
//			c.JSON(http.StatusOK, gin.H{"status": "ok"})
//		}
//	})
//
//	return r
//}

// TODO: move me
const (
	yellowNotFill = "\033[33;20m"
	blueNotFill   = "\033[32;20m"
	green         = "\033[97;42m"
	white         = "\033[90;47m"
	yellow        = "\033[90;43m"
	red           = "\033[97;41m"
	blue          = "\033[97;44m"
	magenta       = "\033[97;45m"
	cyan          = "\033[97;46m"
	reset         = "\033[0m"
)

var (
	_, b, _, _   = runtime.Caller(0)
	basepath     = filepath.Dir(b)
	rootPath     = filepath.Join(basepath, "../../")
	publicAssets = filepath.Join(rootPath, "web/public")
	pathAssets   = filepath.Join(publicAssets, "/assets")
)

func main() {
	log.Println("Go Starting")

	log.Printf("base path -> %s\n", basepath)
	log.Printf("rootPath path -> %s\n", rootPath)
	log.Printf("publicAssets path -> %s\n", publicAssets)
	log.Printf("pathAssets path -> %s\n", pathAssets)

	var appName string = os.Getenv("APP_NAME")
	if appName == "" {
		appName = "httpMaker"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
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

	app := gin.New()
	app.LoadHTMLGlob(filepath.Join(publicAssets, "index.html"))
	app.Static("/assets", pathAssets)
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		var (
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor  = param.ResetColor()
			errorColor  = red
		)
		return fmt.Sprintf("%s%v%s%s|%1s-%1v|%s%s[%3d]%s%s{ %-7s }%s%s=>%s%#v\n%s%s%s",
			yellowNotFill, param.TimeStamp.Format("2006/01/02-15:04:05"), resetColor,
			blueNotFill, param.Latency,
			param.ClientIP, resetColor,

			statusColor, param.StatusCode, resetColor,

			methodColor, param.Method, resetColor,

			cyan, resetColor, param.Path,

			errorColor, param.ErrorMessage, resetColor,
		)
	}))
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Websiste",
		})
	})

	log.Printf("Listening and Serving at 0.0.0.0:%s", port)
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	log.Println("Go Initializing App ... ")

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading the .env file!")
	}

}
