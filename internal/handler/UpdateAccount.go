package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Oriseer/testruns/api"
	"github.com/Oriseer/testruns/internal/tools"

	log "github.com/sirupsen/logrus"
)

type Response struct {
	Username string
	Money    int
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var decoder *json.Decoder = json.NewDecoder(r.Body)
	var database *tools.DatabaseInterface
	var body = Response{}
	var err error

	err = decoder.Decode(&body)

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

	err = (*database).UpdateAccountMoney(body.Username, body.Money)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.Error{
		Message: "Account Updated successfully",
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
