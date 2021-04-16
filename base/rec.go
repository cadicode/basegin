package base

import (
	"runtime"

	"github.com/cadicode/basegin/responseutil"

	"github.com/gin-gonic/gin"
)

// GinRecover define
func GinRecover(c *gin.Context, param interface{}) {
	if err := recover(); err != nil {
		if Logger != nil {
			Logger.WriteError(err, getTrace(), param)
		}
		responseutil.Error(c, "")
	}
}

func getTrace() string {
	var trace [800]byte
	n := runtime.Stack(trace[:], false)
	return string(trace[:n])
}
