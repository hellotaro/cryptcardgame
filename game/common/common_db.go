package common

import (
	"database/sql"
	"fmt"
	"game/repository"
)

func GetCommonDBCon() *sql.DB {
	return repository.GetDBCon("commonDB.db")
}

func InitCommonDB() *sql.DB {
	db := GetCommonDBCon()
	repository.InitTable(db, "Player", `("ID" INTEGER PRIMARY KEY, "Name" VARCHAR(255),"Fund" INTEGER)`)
	repository.InitTable(db, "Game", `("ID" INTEGER PRIMARY KEY, "Name" VARCHAR(255), "ImgUrl" VARCHAR(255), "Description" VARCHAR(255), "Link" VARCHAR(255))`)
	return db
}

func InitPlayers(db *sql.DB) {
	names := []string{"Alice", "Bob", "Cathy", "David", "Emma", "Freddy", "George", "Hass", "Iona", "Jane", "Kathy"}

	for aidx := 0; aidx < len(names); aidx++ {
		InsertPlayer(db, names[aidx], 10000)
	}
}

func InitGames(db *sql.DB) {
	games := []Game{
		{0, "Poker", "https://cdn.pixabay.com/photo/2015/01/08/16/35/play-593207__340.jpg", "wise", "poker"},
		{0, "Blackjack", "https://cdn.pixabay.com/photo/2015/03/20/15/03/poker-682332__340.jpg", "fortune", "blackjack"},
		{0, "Daifugo", "https://cdn.pixabay.com/photo/2022/06/02/11/13/cathedral-7237718__340.jpg", "fun", "daifugo"},
		{0, "Chohan", "https://cdn.pixabay.com/photo/2016/09/08/18/45/cube-1655118__340.jpg", "fun", "chohan"},
		{0, "Othello", "https://cdn.pixabay.com/photo/2016/07/12/11/39/checkmate-1511866__340.jpg", "fun", "othello"},
	}
	for idx := 0; idx < len(games); idx++ {
		_, err := db.Exec(
			`INSERT INTO "Game" (Name, ImgUrl, Description, Link) values (?,?,?,?)`,
			games[idx].Name, games[idx].ImgUrl, games[idx].Description, games[idx].Link,
		)
		if err != nil {
			panic(err)
		}
	}
}

func InsertPlayer(db *sql.DB, name string, fund int) {
	_, err := db.Exec(
		`INSERT INTO "Player" (Name, Fund) values (?,?)`,
		name, fund,
	)
	if err != nil {
		panic(err)
	}
}

func GetAllPlayers(db *sql.DB) []Player {
	rows, err := db.Query(`SELECT * FROM "Player"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	players := make([]Player, 0, 1000)
	for rows.Next() {
		var ID int
		var name string
		var fund int
		if err := rows.Scan(&ID, &name, &fund); err != nil {
			fmt.Println("Error: GetAllPlayers")
		}
		player := Player{ID, name, fund}
		players = append(players, player)
	}
	return players
}

func GetAllGames(db *sql.DB) []Game {
	rows, err := db.Query(`SELECT * FROM "Game"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	games := make([]Game, 0, 1000)
	for rows.Next() {
		var ID int
		var name string
		var imgUrl string
		var description string
		var link string
		if err := rows.Scan(&ID, &name, &imgUrl, &description, &link); err != nil {
			fmt.Println("Error: GetAllGames")
		}
		game := Game{ID, name, imgUrl, description, link}
		games = append(games, game)
	}
	return games
}
