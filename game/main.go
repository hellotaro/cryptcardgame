package main

import (
	"database/sql"
	"fmt"
	"game/poker"
	"game/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InitAccounts(db *sql.DB) {
	names := []string{"Alice", "Bob", "Cathy", "David", "Emma", "Freddy", "George", "Hass", "Iona", "Jane", "Kathy"}

	for aidx := 0; aidx < len(names); aidx++ {
		repository.InsertAccount(db, names[aidx], 10000, "strategy_"+strconv.Itoa(aidx))
	}
	repository.PrintAllAccounts(db)
}

/*
// CUI
func main() {
	//db := poker.InitPokerDB()
	//InitAccounts(db)
	//repository.PrintAllAccounts(db)

	var pokerGame poker.PokerGame
	var userInput string

	gameLoopFlag := true
	userInput = ""
	pokerGame.Phase = 0
	for gameLoopFlag {
		pokerGame = poker.ProceedGameCommand(pokerGame, userInput)

		isSkip := false
		if pokerGame.Phase <= 0 {
			isSkip = true
		}
		if pokerGame.Phase >= 2 && pokerGame.Phase <= 10 {
			if pokerGame.Players[0].State == "dead" {
				isSkip = true
			}
		}

		userInput = ""
		if !isSkip {
			fmt.Scan(&userInput)
		}
	}

}
*/

// web server
func main() {
	//db := poker.InitPokerDB()
	//InitAccounts(db)
	//repository.PrintAllAccounts(db)
	var pokerGame poker.PokerGame

	// web server
	r := gin.Default()
	r.Static("/", "./static")
	r.POST("/api/game/poker/:action", func(c *gin.Context) {
		action := c.Param("action")
		logText := fmt.Sprintf("action: %v\n", action)
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, poker.GameLogInfo{pokerGame.Phase, logText})

		if action == "info" {
			c.JSON(http.StatusOK, pokerGame)
			return
		}
		if action == "table" {
			c.JSON(http.StatusOK, pokerGame)
			return
		}
		if action == "action" {
			var webAction poker.ActionMeta
			c.BindJSON(&webAction)
			command := ""
			if webAction.ActionType == "action" {
				if webAction.Value == "hold" {
					command = "0"
				}
				if webAction.Value == "check" {
					command = "1"
				}
				if webAction.Value == "raise" {
					rfee := webAction.Meta
					command = "2 " + rfee
				}
			}
			if pokerGame.Phase == 1 {
				command = "4"
			}
			logText := fmt.Sprintf("command: %v\n", command)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, poker.GameLogInfo{pokerGame.Phase, logText})
			pokerGame = poker.ProceedGameWeb(pokerGame, command)
			c.JSON(http.StatusOK, pokerGame)
			return
		}
		//c.JSON(http.StatusBadRequest)
	})
	r.Run(":8080")
}
