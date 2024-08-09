package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	log "github.com/sirupsen/logrus"
)

type RequestBody struct {
	Username string
	Token    int
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var database *tools.DatabaseInterface
	var form = RequestBody{}
	var err error

	err = decoder.Decode(&form)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = (*database).DeleteAccount(form.Username, form.Token)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.Error{
		Message: "Account  Deleted Successfully",
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
