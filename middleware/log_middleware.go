package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LogMiddleware struct {
	Handler http.Handler
}

func NewLogMiddleware(handler http.Handler) *LogMiddleware {
	return &LogMiddleware{Handler: handler}
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": request.Method,
		"uri":    request.RequestURI,
		"ip":     request.RemoteAddr,
	}).Infoln("incoming request")

	middleware.Handler.ServeHTTP(writer, request)
}
