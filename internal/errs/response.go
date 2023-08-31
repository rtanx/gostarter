package errs

type Resp struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"error"`
	Appendix any    `json:"appendix_error"`
}
