package middleware

import (
	"errors"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	"github.com/gorilla/schema"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var ErrorMessage = errors.New("\n Invalid Username/Token")
var decoder = schema.NewDecoder()
var username api.UsernameParams

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve Parameters and Header values
		var err error
		var token = r.Header.Get("Authorization")
		var database *tools.DatabaseInterface

		err = decoder.Decode(&username, r.URL.Query())

		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
		}

		if username.Username == "" || token == "" {
			log.Error(ErrorMessage)
			api.RequestErrorHandler(w, ErrorMessage)
			return
		}

		// Connect to database
		database, err = tools.NewDatabase()

		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		// Login details validation
		loginDetails := (*database).GetUserLoginDetails(username.Username)

		if loginDetails == nil || (loginDetails.Token != token) {
			log.Error(ErrorMessage)
			api.RequestErrorHandler(w, ErrorMessage)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := decoder.Decode(&username, r.URL.Query())
		var token = r.Header.Get("Authorization")
		var loginDetails *tools.LoginDetails

		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
		}

		// Connect to DB
		var database *tools.DatabaseInterface

		database, err = tools.NewDatabase()

		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		// Login validation
		loginDetails = (*database).GetUserLoginDetails(username.Username)

		if loginDetails == nil || loginDetails.Token != token {
			log.Error(err)
			api.RequestErrorHandler(w, ErrorMessage)
			return
		}

		next.ServeHTTP(w, r)
	})
}
