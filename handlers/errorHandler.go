package handlers

import "net/http"

type ErrorHandler struct {
}

func (e ErrorHandler) HandleError(writer http.ResponseWriter, req *http.Request) {

}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}
