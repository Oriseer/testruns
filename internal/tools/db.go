package tools

import (
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type LoginDetails struct {
	Username string
	Token    string
}

type MoneyDetails struct {
	Username string
	Money    int
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserMoneyBalance(username string) *MoneyDetails
	SetupDatabase() error
	AddNewAccount(username string, token int) error
	AddAccountMoney(username string, money int) error
	UpdateAccountMoney(username string, newMoney int) error
	DeleteAccount(username string, token int) error
	GetAccount() (*[]MoneyDetails, error)
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	err := database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil

}
