package poker

import (
	"database/sql"
	"fmt"
	"game/cardgame"
	"game/repository"
)

func InitPokerDB() *sql.DB {
	db := repository.GetDBCon("pokerDB.db")
	repository.InitTable(db)
	return db
}
func SetPokerPlayersFromAccount(db *sql.DB, pokerGame *PokerGame) {
	accounts := repository.GetAllAccounts(db)
	for aidx := 0; aidx < pokerGame.PlayerNum; aidx++ {
		pokerGame.Players[aidx].Name = accounts[aidx].Name
		pokerGame.Players[aidx].Strategy = accounts[aidx].Strategy
		pokerGame.Players[aidx].Fund = accounts[aidx].Fund
	}
}

func UpdatePokerPlayersFund(db *sql.DB, players []cardgame.Player) {
	for pidx := 0; pidx < len(players); pidx++ {
		if account, err := repository.GetAccount(db, players[pidx].Name); err == nil {
			repository.UpdateAccountFund(db, account.Name, players[pidx].Fund)
		} else {
			fmt.Println(">Error: Account not found!")
		}
	}
}
