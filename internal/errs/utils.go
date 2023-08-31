package errs

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
	"github.com/rtanx/gostarter/internal/infrastructure/validation"
)

func HandleBindingErr(ctx *gin.Context, err error) {
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		AbortWithHTTPResponse(ctx, ValidationFailed, validation.FailedValidationMapper(verr))
		return
	}
	logger.Error("unexpected error occurred while trying to bind request", logger.Err(err))
	AbortWithHTTPResponse(ctx, BindingFailure, err.Error())

}

func AttchErrToGinCtx(ctx *gin.Context, err APIError, appendix any) {
	ctx.Error(err).SetType(gin.ErrorTypePublic).SetMeta(appendix)
}
