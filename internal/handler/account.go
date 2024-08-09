package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var form *[]tools.MoneyDetails
	var database *tools.DatabaseInterface
	var err error

	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	form, err = (*database).GetAccount()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(form)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
