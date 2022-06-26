package main

import (
	"fmt"
	"game/chinchiro"
	"game/common"
	"game/poker"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	var pokerGame poker.PokerGame
	var chinchiroGame chinchiro.ChinchiroGame
	//db := common.GetCommonDBCon()
	//common.InitCommonDB()
	//common.InitPlayers(db)
	//common.InitGames(db)

	// web server
	r := gin.Default()
	r.Static("/", "./static")
	r.POST("/api/players", func(c *gin.Context) {
		db := common.GetCommonDBCon()
		players := common.GetAllPlayers(db)
		c.JSON(http.StatusOK, players)
	})
	r.POST("/api/games", func(c *gin.Context) {
		db := common.GetCommonDBCon()
		games := common.GetAllGames(db)
		c.JSON(http.StatusOK, games)
	})

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
	r.POST("/api/game/chinchiro/:action", func(c *gin.Context) {
		action := c.Param("action")
		if action == "action" {
			fmt.Println("OKOK")
			chinchiroGame = chinchiro.InitChinchiroGame()
		}
		c.JSON(http.StatusOK, chinchiroGame)
	})

	r.Run(":8080")
}
