package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
)

/* --------------------------------
	App Defaults

Use these in case a .env file is not
configured or when keys in .env are missing
*/
/* --------------------------------*/
const (
	AppNameDefault string = "httpMaker"
	AppPortDefault string = "8080"
)

/*
	 --------------------------------
		App Directories

/* --------------------------------
*/
var (
	_, b, _, _         = runtime.Caller(0)
	Basepath           = filepath.Dir(b)
	RootPath           = filepath.Join(Basepath, "../../../")
	PublicAssets       = filepath.Join(RootPath, "web/public")
	PathIndex          = filepath.Join(PublicAssets, "index.html")
	PathAssetsRelative = "/assets"
	PathAssets         = filepath.Join(PublicAssets, PathAssetsRelative)
)

/* --------------------------------
	LOGGER
/* --------------------------------*/

// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
// By default gin.DefaultWriter = os.Stdout
func LoggerWithFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(appLogger)
}

func appLogger(param gin.LogFormatterParams) string {
	// logger configuration
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
	// custom formatters
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
}
