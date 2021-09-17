package exception

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if alreadyExistError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		// Create slice of error
		var errors []string
		var customError string

		// Custom error
		for _, e := range exception {
			switch e.Tag() {
			case "required":
				customError = fmt.Sprintf("%s is required", e.Field())
			case "min":
				customError = fmt.Sprintf("%s is minimum %s character", e.Field(), e.Param())
			case "max":
				customError = fmt.Sprintf("%s is maximum %s character", e.Field(), e.Param())
			default:
				customError = e.Error()
			}
			errors = append(errors, customError)
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   errors,
		}

		helper.WriteToResponseBody(writer, webResponse)
		log.WithFields(log.Fields{
			"code":   http.StatusBadRequest,
			"errors": errors,
			"at":     time.Now().Format("2006-01-02 15:04:05"),
			"method": request.Method,
			"uri":    request.RequestURI,
			"ip":     request.RemoteAddr,
		}).Errorln("validation error")

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		log.WithFields(log.Fields{
			"code":   http.StatusNotFound,
			"error":  exception.Error,
			"at":     time.Now().Format("2006-01-02 15:04:05"),
			"method": request.Method,
			"uri":    request.RequestURI,
			"ip":     request.RemoteAddr,
		}).Errorln("not found error")

		return true
	} else {
		return false
	}
}

func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AlreadyExistError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		webResponse := web.WebResponse{
			Code:   http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		log.WithFields(log.Fields{
			"code":   http.StatusConflict,
			"error":  exception.Error,
			"at":     time.Now().Format("2006-01-02 15:04:05"),
			"method": request.Method,
			"uri":    request.RequestURI,
			"ip":     request.RemoteAddr,
		}).Errorln("already exist error")

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
	log.WithFields(log.Fields{
		"code":   http.StatusInternalServerError,
		"error":  err,
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": request.Method,
		"uri":    request.RequestURI,
		"ip":     request.RemoteAddr,
	}).Errorln("internal server error")
}
