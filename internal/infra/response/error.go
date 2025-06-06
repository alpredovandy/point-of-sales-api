package response

import "net/http"

func newError(httpStatusCode int, msg string, err error, opts ...OptResponse) Response {
	var resp = Response{
		HttpStatus: httpStatusCode,
		Message:    msg,
		Success:    false,
		Error:      err.Error(),
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewErrorGeneral(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusInternalServerError, msg, err, opts...)
}

func NewErrorBadRequest(msg string, err error, opts ...OptResponse) Response {
	return newError(http.StatusBadRequest, msg, err, opts...)
}
