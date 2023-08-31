package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rtanx/gostarter/internal/errs"
)

// ErrorReporter usefull when we want report the error to 3rdparty like sentry
// and we will define it here. Currently not used
func ErrorReporter() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.ByType(gin.ErrorTypePublic).Last()

		if err != nil {
			fmt.Println(err.Err.(errs.APIError))
			if apierr, ok := err.Err.(errs.APIError); ok {
				errs.AbortWithHTTPResponse(ctx, apierr, err.Meta)
				return
			}
			errs.AbortWithHTTPResponse(ctx, errs.InternalServer, nil)
			return
		}
	}
}
