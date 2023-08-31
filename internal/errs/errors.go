package errs

import "net/http"

var (
	InvalidParam         APIError = newAPIError(http.StatusBadRequest, "invalid parameter")
	ValidationFailed     APIError = newAPIError(http.StatusBadRequest, "some fields in the request do not meet the criteria")
	InternalServer       APIError = newAPIError(http.StatusInternalServerError, "internal server error")
	Unauthorized         APIError = newAPIError(http.StatusUnauthorized, "unauthorized")
	Unauthenticated      APIError = newAPIError(http.StatusUnauthorized, "unauthenticated")
	AccessDenied         APIError = newAPIError(http.StatusForbidden, "you have no right to access intended resources")
	UnprocessableUser    APIError = newAPIError(http.StatusUnprocessableEntity, "failed to process user entities or resources")
	BearerTokenRequired  APIError = newAPIError(http.StatusBadRequest, "bearer token not provided")
	BearerTokenInvalid   APIError = newAPIError(http.StatusUnauthorized, "invalid or expired bearer token")
	WrongCredentials     APIError = newAPIError(http.StatusUnauthorized, "wrong credentials")
	InvalidRefresToken   APIError = newAPIError(http.StatusBadRequest, "invalid refresh token")
	BindingFailure       APIError = newAPIError(http.StatusBadRequest, "failed to bind the request")
	UserNotFound         APIError = newAPIError(http.StatusBadRequest, "user not found")
	InvalidOTP           APIError = newAPIError(http.StatusBadRequest, "invalid or expired otp code")
	InvalidResetPwdToken APIError = newAPIError(http.StatusBadRequest, "invalid or expired reset password token")
	EmptyResult          APIError = newAPIError(http.StatusNotFound, "empty records")
	ResultNotFound       APIError = newAPIError(http.StatusNotFound, "record not found")
)
