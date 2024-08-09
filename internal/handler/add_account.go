package handler

import (
	"encoding/json"

	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	log "github.com/sirupsen/logrus"
)

func AddAccount(w http.ResponseWriter, r *http.Request) {
	var accountForm = api.AddNewAccountForm{}
	var decoder = json.NewDecoder(r.Body)
	var database *tools.DatabaseInterface

	err := decoder.Decode(&accountForm)

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

	err = (*database).AddNewAccount(accountForm.Username, accountForm.Token)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.Error{
		Message: "Account successfully inserted",
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
