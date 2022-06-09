package poker

import (
	"game/cardgame"
	"os"
	"strconv"
	"strings"
)

type ActionMeta struct {
	ActionType string
	Value      string
	Meta       string
}

func ProceedGameWeb(pokerGame PokerGame, command string) PokerGame {
	// ゲーム作成
	if pokerGame.Phase == 1 {
		playerNum, _ := strconv.Atoi(command)
		db := InitPokerDB()
		newPokerGame := InitPokerGame(playerNum)
		SetPokerPlayersFromAccount(db, &newPokerGame)
		newPokerGame.Phase = pokerGame.Phase
		pokerGame = newPokerGame
		pokerGame.Fee = pokerGame.FirstFee
		CalcPokerGameScore(&pokerGame)
	}
	// 1回目アクション
	if pokerGame.Phase == 2 {
		pokerGame.Players[0].Info = "My Action: " + command
		cmd_detail := strings.Split(command, " ")
		actionNum, _ := strconv.Atoi(cmd_detail[0])
		if len(cmd_detail) == 1 {
			MEAction(&pokerGame, actionNum)
		}
		if len(cmd_detail) == 2 {
			raisedFee, _ := strconv.Atoi(cmd_detail[1])
			MEActionRaise(&pokerGame, raisedFee)
		}
	}
	// 1回目NPCアクション
	if pokerGame.Phase == 4 {
		//NPC
		isRaised := AllNPCAction(&pokerGame)
		if isRaised {
			pokerGame.Phase -= (2 + 1)
		} else {
			cardgame.AddCardToTable(&pokerGame.Table, &pokerGame.Deck)
			EndGameTurn(&pokerGame)
			CalcPokerGameScore(&pokerGame)
		}
	}
	// 2回目アクション
	if pokerGame.Phase == 5 {
		pokerGame.Players[0].Info = "My Action: " + command
		cmd_detail := strings.Split(command, " ")
		actionNum, _ := strconv.Atoi(cmd_detail[0])
		if len(cmd_detail) == 1 {
			MEAction(&pokerGame, actionNum)
		}
		if len(cmd_detail) == 2 {
			raisedFee, _ := strconv.Atoi(cmd_detail[1])
			MEActionRaise(&pokerGame, raisedFee)
		}
	}
	// 2回目NPCアクション
	if pokerGame.Phase == 7 {
		//NPC
		isRaised := AllNPCAction(&pokerGame)
		if isRaised {
			pokerGame.Phase -= (2 + 1)
		} else {
			cardgame.AddCardToTable(&pokerGame.Table, &pokerGame.Deck)
			EndGameTurn(&pokerGame)
			CalcPokerGameScore(&pokerGame)
		}
	}
	// 3回目アクション
	if pokerGame.Phase == 8 {
		pokerGame.Players[0].Info = "My Action: " + command
		cmd_detail := strings.Split(command, " ")
		actionNum, _ := strconv.Atoi(cmd_detail[0])
		if len(cmd_detail) == 1 {
			MEAction(&pokerGame, actionNum)
		}
		if len(cmd_detail) == 2 {
			raisedFee, _ := strconv.Atoi(cmd_detail[1])
			MEActionRaise(&pokerGame, raisedFee)
		}
	}
	// 3回目NPCアクション
	if pokerGame.Phase == 10 {
		//NPC
		isRaised := AllNPCAction(&pokerGame)
		if isRaised {
			pokerGame.Phase -= (2 + 1)
		} else {
			EndGameTurn(&pokerGame)
			CalcPokerGameScore(&pokerGame)
			rankedPlayers := cardgame.SetPlayerRankingIndex(pokerGame.Players)

			DistributePod(&pokerGame, rankedPlayers[0])
		}
	}
	if pokerGame.Phase == 11 {
		actionNum, _ := strconv.Atoi(command)
		if actionNum == 1 {
			pokerGame.Phase = 0
		} else {
			os.Exit(0)
		}
	}

	pokerGame.Phase += 1

	return pokerGame
}
