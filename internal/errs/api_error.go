package errs

type APIError interface {
	error
	StatusCode() int
}
type apiError struct {
	Code int    `json:"code"`
	Msg  string `json:"error"`
}

func (ae *apiError) StatusCode() int {
	return ae.Code
}

func (ae *apiError) Error() string {
	return ae.Msg
}

func newAPIError(code int, errMsg string) *apiError {
	return &apiError{
		Code: code,
		Msg:  errMsg,
	}
}

// New is shortcut and exported newAPIError
func New(code int, errMsg string) *apiError {
	return &apiError{
		Code: code,
		Msg:  errMsg,
	}
}
