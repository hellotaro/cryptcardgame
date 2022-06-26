package poker

import (
	"database/sql"
	"errors"
	"fmt"
	"game/cardgame"
	"game/repository"
	"strconv"
)

type Account struct {
	ID       int
	Name     string
	Fund     int
	Strategy string
}

func InitPokerDB() *sql.DB {
	db := repository.GetDBCon("pokerDB.db")
	repository.InitTable(db, "Account", `("ID" INTEGER PRIMARY KEY, "Name" VARCHAR(255),"Fund" INTEGER,"Strategy" VARCHAR(255))`)
	return db
}
func SetPokerPlayersFromAccount(db *sql.DB, pokerGame *PokerGame) {
	accounts := GetAllAccounts(db)
	for aidx := 0; aidx < pokerGame.PlayerNum; aidx++ {
		pokerGame.Players[aidx].Name = accounts[aidx].Name
		pokerGame.Players[aidx].Strategy = accounts[aidx].Strategy
		pokerGame.Players[aidx].Fund = accounts[aidx].Fund
	}
}

func UpdatePokerPlayersFund(db *sql.DB, players []cardgame.Player) {
	for pidx := 0; pidx < len(players); pidx++ {
		if account, err := GetAccount(db, players[pidx].Name); err == nil {
			UpdateAccountFund(db, account.Name, players[pidx].Fund)
		} else {
			fmt.Println(">Error: Account not found!")
		}
	}
}

func InsertAccount(db *sql.DB, name string, fund int, strategy string) {
	_, err := db.Exec(
		`INSERT INTO "Account" (Name, Fund, Strategy) values (?,?,?)`,
		name, fund, strategy,
	)
	if err != nil {
		panic(err)
	}
}

func UpdateAccountFund(db *sql.DB, name string, fund int) {
	_, err := db.Exec(
		`UPDATE ACCOUNT SET Fund=? WHERE Name=?`,
		fund, name,
	)
	if err != nil {
		panic(err)
	}
}

func UpdateAccountStrategy(db *sql.DB, name string, strategy string) {
	_, err := db.Exec(
		`UPDATE ACCOUNT SET Strategy=? WHERE Name=?`,
		strategy, name,
	)
	if err != nil {
		panic(err)
	}
}

func DeleteAccount(db *sql.DB, name string) {
	_, err := db.Exec(
		`DELETE FROM ACCOUNT WHERE Name=?`,
		name,
	)
	if err != nil {
		panic(err)
	}
}

func GetAllAccounts(db *sql.DB) []Account {
	rows, err := db.Query(`SELECT * FROM "Account"`)
	if err != nil {
		panic(err)
	}

	accounts := make([]Account, 0, 1000)

	defer rows.Close()
	for rows.Next() {
		var ID int
		var name string
		var fund int
		var strategy string

		if err := rows.Scan(&ID, &name, &fund, &strategy); err != nil {
			fmt.Println("Error: getAllAccounts")
		}

		account := Account{ID, name, fund, strategy}
		accounts = append(accounts, account)
	}

	return accounts
}

func GetAccount(db *sql.DB, account_name string) (Account, error) {
	row := db.QueryRow(`SELECT * FROM "Account" WHERE Name = ?`, account_name)

	var ID int
	var name string
	var fund int
	var strategy string

	err := row.Scan(&ID, &name, &fund, &strategy)

	var account Account
	var reserr error
	reserr = nil
	switch {
	case err == sql.ErrNoRows:
		account = Account{}
		reserr = errors.New("no rows")
	case err != nil:
		reserr = errors.New("db error")
		panic(err)
	default:
		account = Account{ID, name, fund, strategy}
	}

	return account, reserr
}

func PrintAllAccounts(db *sql.DB) {
	accounts := GetAllAccounts(db)
	for idx := 0; idx < len(accounts); idx++ {
		fmt.Println(strconv.Itoa(accounts[idx].ID) + " , " + accounts[idx].Name + " , " + strconv.Itoa(accounts[idx].Fund) + " , " + accounts[idx].Strategy)
	}
}
