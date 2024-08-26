package middleware

import (
	"errors"
	"net/http"

	"github.com/milanmlft/goapi/api"
	"github.com/milanmlft/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorisedError = errors.New("Invalid username or token.")

// Middleware authorisation function
func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Authorisation")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return
		}

		// Set up database connection
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return
		}

		// Call next middleware or handler funtion in line
		next.ServeHTTP(w, r)
	})
}
