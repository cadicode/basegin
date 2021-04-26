package basegin

import (
	"errors"

	"github.com/cadicode/basegin/base"
	"github.com/gin-contrib/cors"

	"github.com/cadicode/basegin/router"
	"github.com/gin-gonic/gin"
)

var (
	corsAllowAllOrigins = []string{"*"}
)

// CreateDefaultGin Create *gin.Engine with default log,using logrus
func CreateDefaultGin(
	isProduct bool,
	isCors bool,
	logger base.ILogger,
	corsAllowOrigins []string,
	corsAllowHeaders []string,
	groupedControllers map[string][]router.IBaseController) (*gin.Engine, error) {

	if len(logFolderPath) == 0 {
		return nil, errors.New("log folder path is nil")
	}
	logger, err := base.NewLogrusLogger(logFolderPath)
	if err != nil {
		return nil, err
	}
	return CreateGin(isProduct, isCors, logger, corsAllowAllOrigins, corsAllowHeaders, groupedControllers)
}

// CreateGin create *gin.Engine with custom logger
func CreateGin(
	isProduct bool,
	isCors bool,
	logger base.ILogger,
	corsAllowOrigins []string,
	corsAllowHeaders []string,
	groupedControllers map[string][]router.IBaseController) (*gin.Engine, error) {

	setMode(isProduct)

	r := gin.Default()

	if logger != nil {
		base.SetLogger(logger)
	}

	if isCors {
		setCors(r, corsAllowOrigins, corsAllowHeaders)
	}

	r.RedirectFixedPath = true

	for k, v := range groupedControllers {
		router.RegisterGroupAPIRoute(k, r, v)
	}
	return r, nil
}

func setMode(isProduct bool) {
	if isProduct {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

// func hookLogger(logFolderPath string) (*base.LogrusLogger, error) {
// 	return base.NewLogrusLogger(logFolderPath)
// }

func setCors(r *gin.Engine, corsAllowOrigins []string, allowHeaders []string) {
	if len(corsAllowOrigins) == 0 {
		corsAllowOrigins = corsAllowAllOrigins
	}
	c := cors.DefaultConfig()
	c.AllowOrigins = corsAllowOrigins
	c.AllowCredentials = true
	if len(allowHeaders) > 0 {
		c.AddAllowHeaders(allowHeaders...)
	}
	r.Use(cors.New(c))
}

// CreateGin2 create gin.Engine without mapping controllers
func CreateGin2(
	isProduct bool,
	isCors bool,
	logFolderPath string,
	corsAllowOrigins []string,
	corsAllowHeaders []string) (*gin.Engine, error) {
	setMode(isProduct)

	r := gin.Default()

	if logger != nil {
		base.SetLogger(logger)
	}

	if isCors {
		setCors(r, corsAllowOrigins, corsAllowHeaders)
	}

	r.RedirectFixedPath = true
	return r, nil
}
