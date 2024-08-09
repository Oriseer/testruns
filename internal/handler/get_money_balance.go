package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetMoneyBalance(w http.ResponseWriter, r *http.Request) {
	// Retrieve the Request Parameters
	var decoder *schema.Decoder = schema.NewDecoder()
	var username api.UsernameParams
	var err error
	var database *tools.DatabaseInterface

	err = decoder.Decode(
		&username,
		r.URL.Query(),
	)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	// Connect to Database
	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Retrieve Money Balance and create response
	var moneybalance = (*database).GetUserMoneyBalance(username.Username)

	var response = api.MoneyBalanceResponse{
		MoneyBalance: moneybalance.Money,
		Code:         http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
