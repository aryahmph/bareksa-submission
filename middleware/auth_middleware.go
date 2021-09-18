package middleware

import (
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/model/web"
	"net/http"
	"regexp"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "BAREKSA_INTERNSHIP" == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		regex, err := regexp.Compile(`/uploads/news/(.*)`)
		helper.PanicIfError(err)
		if regex.MatchString(request.RequestURI) {
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	}
}
