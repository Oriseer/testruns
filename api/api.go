package api

import (
	"encoding/json"
	"net/http"
)

type AddNewAccountForm struct {
	Username string
	Token    int
}

type AddAccountMoneyForm struct {
	Username string
	Money    int
}

type UsernameParams struct {
	Username string
}

type MoneyBalanceResponse struct {
	MoneyBalance int
	Code         int
}

type Error struct {
	Message string
	Code    int
}

func writeError(w http.ResponseWriter, message string, code int) {
	var response = Error{
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Error occured", http.StatusInternalServerError)
	}
)
