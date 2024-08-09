package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	log "github.com/sirupsen/logrus"
)

func AddUserMoney(w http.ResponseWriter, r *http.Request) {
	var decoder *json.Decoder = json.NewDecoder(r.Body)
	var accountMoneyForm = api.AddAccountMoneyForm{}
	var database *tools.DatabaseInterface

	err := decoder.Decode(&accountMoneyForm)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Connect to DB
	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = (*database).AddAccountMoney(accountMoneyForm.Username, accountMoneyForm.Money)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.Error{
		Message: `Account money successfully inserted`,
		Code:    http.StatusOK,
	}

	w.Header().Set(
		"Content-type",
		`application/json`,
	)
	json.NewEncoder(w).Encode(response)

}
