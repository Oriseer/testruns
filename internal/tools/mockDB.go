package tools

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Khel": {
		Username: "Khel",
		Token:    "123GWTR",
	},

	"Emz": {
		Username: "Emz",
		Token:    "EMZ310",
	},
}

var mockMoneyBalance = map[string]MoneyDetails{
	"Khel": {
		Username: "Khel",
		Money:    90,
	},

	"Emz": {
		Username: "Emz",
		Money:    80,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	var loginDetails = LoginDetails{}

	//loginDetails, ok := mockLoginDetails[username]

	db, err := d.connectDB()

	if err != nil {
		log.Error(err)
		return nil
	}

	err = db.Get(&loginDetails, "SELECT * FROM LOGIN_DETAILS WHERE USERNAME=$1", username)

	if err != nil {
		log.Error(err)
		return nil
	}

	return &loginDetails
}

func (d *mockDB) GetUserMoneyBalance(username string) *MoneyDetails {
	var userBalancedetails = MoneyDetails{}

	//userBalancedetails = mockMoneyBalance[username]

	db, err := d.connectDB()

	if err != nil {
		log.Error(err)
		return nil
	}

	err = db.Get(&userBalancedetails, "SELECT * FROM MONEY_DETAILS WHERE USERNAME=$1", username)

	if err != nil {
		log.Error(err)
		return nil
	}

	return &userBalancedetails
}

func (d *mockDB) AddNewAccount(username string, token int) error {
	db, err := d.connectDB()
	var insert_statement string = "INSERT INTO LOGIN_DETAILS(username, token) VALUES($1, $2)"

	if err != nil {
		log.Error(err)
		return err
	}

	tx := db.MustBegin()

	tx.MustExec(insert_statement, username, token)
	tx.Commit()

	return nil
}

func (d *mockDB) AddAccountMoney(username string, money int) error {
	db, err := d.connectDB()
	var insert_statement string = "INSERT INTO MONEY_DETAILS(username, money) VALUES($1, $2)"

	if err != nil {
		log.Error(err)
		return err
	}

	tx := db.MustBegin()

	tx.MustExec(insert_statement, username, money)
	tx.Commit()

	return nil
}

func (d *mockDB) UpdateAccountMoney(username string, newToken int) error {
	db, err := d.connectDB()
	var update_statement string = "UPDATE MONEY_DETAILS SET MONEY = $1 WHERE USERNAME = $2"

	if err != nil {
		log.Error(err)
		return err
	}

	tx := db.MustBegin()
	tx.MustExec(update_statement, newToken, username)
	tx.Commit()

	return nil
}

func (d *mockDB) DeleteAccount(username string, token int) error {
	var delete_statement string = "DELETE FROM LOGIN_DETAILS WHERE USERNAME = $1 AND TOKEN = $2"

	db, err := d.connectDB()

	if err != nil {
		log.Error(err)
		return err
	}

	tx := db.MustBegin()
	tx.MustExec(delete_statement, username, token)
	tx.Commit()

	return nil
}

func (d *mockDB) GetAccount() (*[]MoneyDetails, error) {
	db, err := d.connectDB()
	var moneyDetails = []MoneyDetails{}
	var query_statement string = "SELECT * FROM MONEY_DETAILS"

	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = db.Select(&moneyDetails, query_statement)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &moneyDetails, nil

}

func (d *mockDB) SetupDatabase() error {

	_, err := d.connectDB()
	if err != nil {
		return err
	}
	return nil
}

func (d *mockDB) connectDB() (*sqlx.DB, error) {
	var db_details string = "user=postgres dbname=testruns sslmode=disable password=testpost host=localhost"
	db, err := sqlx.Connect(
		"postgres",
		db_details,
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, err
}
