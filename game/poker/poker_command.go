package poker

import (
	"fmt"
	"game/cardgame"
	"os"
	"strconv"
	"strings"
)

func PrintActionList() {
	fmt.Println("アクションを選択してください。")
	fmt.Println("  0: Hold")
	fmt.Println("  1: Check")
	fmt.Println("  2: Raise")
}
func PrintRaise(pokerGame PokerGame) {
	fmt.Println("レイズベットを記入してください。")
	fmt.Printf("　　Fee: %v, 現在手持ち：%v(-%v)\n", pokerGame.Fee, (pokerGame.Players[0].Fund - pokerGame.Players[0].Bet), pokerGame.Players[0].Bet)
}
func RaiseCheck(pokerGame *PokerGame, betNum int) bool {
	if betNum <= 0 {
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "正しい値を入力してください！"})
		pokerGame.Phase -= 1
		return false
	}
	if betNum+pokerGame.Players[0].Bet <= pokerGame.Fee {
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "Feeよりも大きいベットを入力してください！"})
		pokerGame.Phase -= 1
		return false
	}

	if betNum <= pokerGame.Players[0].Fund {
		pokerGame.Players[0].Bet = betNum
		pokerGame.Fee = pokerGame.Players[0].Bet
		pokerGame.LastRaisedPlayerIdx = 0
		pokerGame.Phase += 2
		return true
	} else {
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "ベットに対して手持ちが不足しています！"})
		pokerGame.Phase -= 1
		return false
	}
}
func PrintSituation(pokerGame PokerGame) {
	cardgame.PrintTable(pokerGame.Table)
	meplayer := pokerGame.Players[0]
	PrintPokerPlayer(meplayer, pokerGame.Table)
	cardgame.PrintPlayersBet(pokerGame.Players)
	nextscores := CalcNextHandsetScores(append(pokerGame.Table.Cards, meplayer.Hand.Cards...), 1)
	fmt.Printf("Prob: %v\n", CalcNextHandsetProb(nextscores, "pair"))
	/*
		nextscores = FilterNextHandsetScores(nextscores, "pair")
		fmt.Printf("[NextHand]\n")
		for idx := 0; idx < len(nextscores); idx++ {
			fmt.Printf("{%v %v}\n", cardgame.GetCardsStr(nextscores[idx].Cards), EncodeHandset(nextscores[idx].Score))
		}
	*/

	PrintActionList()
}

func MEAction(pokerGame *PokerGame, actionNum int) {
	logText := fmt.Sprintf("MEAction: %v\n", actionNum)
	pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})

	if pokerGame.Players[0].State == "live" {
		if actionNum == 0 {
			pokerGame.Players[0].State = "dead"
		}
		if actionNum == 1 {
			if pokerGame.Players[0].Fund >= pokerGame.Fee {
				pokerGame.Players[0].Bet = pokerGame.Fee
			} else {
				// all_in
				pokerGame.Players[0].Bet = pokerGame.Players[0].Fund
				pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "ALL IN!!!!!!"})
			}
		}
	}
	pokerGame.Phase += 2
}

func MEActionRaise(pokerGame *PokerGame, fee int) {
	if pokerGame.Players[0].Fund-pokerGame.Players[0].Bet == 0 {
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "すでにオールインしています！"})
		pokerGame.Phase += 2
	} else {
		RaiseCheck(pokerGame, fee)
	}
}

func NPCAction(pokerGame *PokerGame, pidx int, action PlayAction) cardgame.Player {
	logText := fmt.Sprintf("NPCAction: %v\n", pidx)
	pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})

	player := pokerGame.Players[pidx]

	if player.State == "dead" {
		return player
	}

	if action.ActionType == "check" {
		player.Bet = action.Cost
	}
	if action.ActionType == "raise" || action.ActionType == "all_in" {
		player.Bet = action.Cost
		if action.Cost > 0 {
			pokerGame.LastRaisedPlayerIdx = pidx
		}
	}
	if action.ActionType == "hold" {
		player.State = "dead"
	}

	return player
}
func AllNPCAction(pokerGame *PokerGame) bool {
	pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "AllNPCAction"})

	isRaised := false
	for pidx := 1; pidx < pokerGame.PlayerNum; pidx++ {
		if pokerGame.LastRaisedPlayerIdx >= 1 && pokerGame.LastRaisedPlayerIdx == pidx {
			break
		}

		action := NPCPlay(*pokerGame, pidx, pokerGame.Fee)
		pokerGame.Players[pidx] = NPCAction(pokerGame, pidx, action)
		pokerGame.GameHistory = append(pokerGame.GameHistory, action)
		if action.ActionType == "raise" {
			isRaised = true
			pokerGame.Fee = action.Cost
			logText := fmt.Sprintf("%v has Raised!: Fee: %v\n", pokerGame.Players[pidx].Name, pokerGame.Fee)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			pokerGame.Players[pidx].Info = logText
		}
		if action.ActionType == "all_in" {
			if action.Cost > 0 {
				isRaised = true
				if pokerGame.Fee < action.Cost {
					pokerGame.Fee = action.Cost
				}
				logText := fmt.Sprintf("%v has AllIn!!!\n", pokerGame.Players[pidx].Name)
				pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
				pokerGame.Players[pidx].Info = logText
			}
		}
		if action.ActionType == "check" {
			logText := fmt.Sprintf("%v has Checked!\n", pokerGame.Players[pidx].Name)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			pokerGame.Players[pidx].Info = logText
		}
		if action.ActionType == "hold" {
			logText := fmt.Sprintf("%v has Held!\n", pokerGame.Players[pidx].Name)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			pokerGame.Players[pidx].Info = logText
		}
	}
	return isRaised
}

func EndGameTurn(pokerGame *PokerGame) {
	pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "EndGameTurn: %v\n"})

	pokerGame.Fee = 0
	pokerGame.LastRaisedPlayerIdx = -1
	for pidx := 0; pidx < len(pokerGame.Players); pidx++ {
		pokerGame.Table.Pod += pokerGame.Players[pidx].Bet
		pokerGame.Players[pidx].Fund -= pokerGame.Players[pidx].Bet
		pokerGame.Players[pidx].Bet = 0
		pokerGame.Players[pidx].Info = ""
	}
}

func DistributePod(pokerGame *PokerGame, player cardgame.Player) {
	for pidx := 0; pidx < len(pokerGame.Players); pidx++ {
		if pokerGame.Players[pidx].Name == player.Name {
			pokerGame.Players[pidx].Fund += pokerGame.Table.Pod
			pokerGame.Table.Pod = 0
			break
		}
	}
}

func ProceedGameCommand(pokerGame PokerGame, command string) PokerGame {
	logText := fmt.Sprintf("ProceedGameCommand: %v\n", pokerGame.Phase)
	pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})

	if pokerGame.Phase == 0 {
		pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, "プレーヤーの人数を入力してください。"})
	}
	// ゲーム作成
	if pokerGame.Phase == 1 {
		playerNum, _ := strconv.Atoi(command)
		db := InitPokerDB()
		newPokerGame := InitPokerGame(playerNum)
		SetPokerPlayersFromAccount(db, &newPokerGame)
		newPokerGame.Phase = pokerGame.Phase
		pokerGame = newPokerGame
		pokerGame.Fee = pokerGame.FirstFee

		PrintSituation(pokerGame)
	}
	// 1回目アクション
	if pokerGame.Phase == 2 {
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
			logText := fmt.Sprintf("Fee: %v\n", pokerGame.Fee)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			cardgame.PrintPlayersBet(pokerGame.Players)
			PrintActionList()
		} else {
			cardgame.AddCardToTable(&pokerGame.Table, &pokerGame.Deck)
			EndGameTurn(&pokerGame)
			PrintSituation(pokerGame)
		}
	}
	// 2回目アクション
	if pokerGame.Phase == 5 {
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
			logText := fmt.Sprintf("Fee: %v\n", pokerGame.Fee)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			cardgame.PrintPlayersBet(pokerGame.Players)
			PrintActionList()
		} else {
			cardgame.AddCardToTable(&pokerGame.Table, &pokerGame.Deck)
			EndGameTurn(&pokerGame)
			PrintSituation(pokerGame)
		}
	}
	// 3回目アクション
	if pokerGame.Phase == 8 {
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
			logText := fmt.Sprintf("Fee: %v\n", pokerGame.Fee)
			pokerGame.GameLogInfos = append(pokerGame.GameLogInfos, GameLogInfo{pokerGame.Phase, logText})
			cardgame.PrintPlayersBet(pokerGame.Players)
			PrintActionList()
		} else {
			cardgame.AddCardToTable(&pokerGame.Table, &pokerGame.Deck)
			EndGameTurn(&pokerGame)
			PrintSituation(pokerGame)

			CalcPokerGameScore(&pokerGame)

			fmt.Println("===[Ranking]===")
			rankedPlayers := cardgame.SetPlayerRankingIndex(pokerGame.Players)
			DistributePod(&pokerGame, rankedPlayers[0])
			PrintPokerPlayers(rankedPlayers, pokerGame.Table)

			fmt.Println("もう一度ゲームをプレイしますか？")
			fmt.Println("  0: いいえ")
			fmt.Println("  1: はい")
		}
	}
	if pokerGame.Phase == 11 {
		actionNum, _ := strconv.Atoi(command)
		if actionNum == 1 {
			pokerGame.Phase = -1
		} else {
			os.Exit(0)
		}
	}

	pokerGame.Phase += 1

	return pokerGame
}
