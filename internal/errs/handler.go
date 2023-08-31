package errs

import (
	"github.com/gin-gonic/gin"
)

func AbortWithHTTPResponse(ctx *gin.Context, err error, appendixErr any) {
	var (
		statusCode int
		errMsg     string
	)
	switch e := err.(type) {
	case APIError:
		statusCode = e.StatusCode()
		errMsg = e.Error()
	default:
		statusCode = InternalServer.StatusCode()
		errMsg = InternalServer.Error()
	}
	ctx.AbortWithStatusJSON(statusCode, Resp{
		Code:     statusCode,
		ErrorMsg: errMsg,
		Appendix: appendixErr,
	})
}

func IsAPIError(err error) bool {
	if _, ok := err.(APIError); ok {
		return true
	}
	return false
}
