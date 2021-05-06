package responseutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseObj define  a response object
type ResponseObj struct {
	Code ResponseCode `json:"code"`
	Data interface{}  `json:"data"`
}

// ResponseCode define the new type for response code
type ResponseCode int32

const (
	// RSuccess is success status
	RSuccess ResponseCode = 200
	// Client Error
	RClientError ResponseCode = 400
	// Server Error
	RServerError ResponseCode = 500
	// ROther is other status
	ROther
)

// GinResponseObj convert ResponseObj into gin.H
func GinResponseObj(o *ResponseObj) gin.H {
	return gin.H{
		"code": o.Code,
		"data": o.Data,
	}
}

// Error response error message
func Error(c *gin.Context, additionalInfo string) {
	ErrorWithCode(c, RServerError, additionalInfo)
}

func ErrorWithCode(c *gin.Context, code ResponseCode, msg string) {
	if len(msg) == 0 {
		msg = "system error occurred"
	}
	result := ResponseObj{
		Code: code,
		Data: msg,
	}
	c.JSON(http.StatusOK, GinResponseObj(&result))
}

func Success(c *gin.Context, data interface{}) {
	result := ResponseObj{
		RSuccess,
		data,
	}
	c.JSON(http.StatusOK, GinResponseObj(&result))
}
