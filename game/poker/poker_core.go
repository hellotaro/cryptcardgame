package poker

import (
	"fmt"
	"game/cardgame"
	"strconv"
)

type PokerGame struct {
	Players             []cardgame.Player
	PlayerNum           int
	Deck                cardgame.Deck
	Table               cardgame.Table
	Phase               int
	Fee                 int
	FirstFee            int
	LastRaisedPlayerIdx int
	GameHistory         []PlayAction
	GameLogInfos        []GameLogInfo
	GameInfo            string
}

type HandsetInfo struct {
	Yaku   string
	Number int
}

type NextCardScoreInfo struct {
	Cards []cardgame.Card
	Score int
}

type GameLogInfo struct {
	Phase int
	Text  string
}

func InitPlayers(deck *cardgame.Deck, playerNum int) []cardgame.Player {
	players := make([]cardgame.Player, playerNum)
	for pidx := 0; pidx < playerNum; pidx++ {
		players[pidx].Name = "player-" + strconv.Itoa(pidx)
		hand := cardgame.Hand{}
		hand.Cards = make([]cardgame.Card, 2)
		for cidx := 0; cidx < 2; cidx++ {
			hand.Cards[cidx] = cardgame.DrawCardFromDeck(deck)
		}
		players[pidx].Hand = hand
		players[pidx].Strategy = "strategy"
		players[pidx].State = "live"
		players[pidx].Score = 0
		players[pidx].Rank = -1
		players[pidx].Bet = 0
	}

	return players
}

func PrintPokerPlayer(player cardgame.Player, table cardgame.Table) {
	cardgame.PrintPlayer(player, table)
	score := GetHandScore(append(table.Cards, player.Hand.Cards...))
	fmt.Printf("[Yaku] ")
	handset := EncodeHandset(score)
	for hidx := 0; hidx < len(handset); hidx++ {
		fmt.Printf("%v ", handset[hidx])
	}
	fmt.Printf("\n")
}

func PrintPokerPlayers(players []cardgame.Player, table cardgame.Table) {
	fmt.Println("===Players===")
	for pidx := 0; pidx < len(players); pidx++ {
		PrintPokerPlayer(players[pidx], table)
	}
}

func InitPokerGame(playerNum int) PokerGame {
	pokerGame := PokerGame{}
	pokerGame.Phase = 0
	pokerGame.PlayerNum = playerNum
	pokerGame.Deck = cardgame.InitDeck()
	pokerGame.Players = InitPlayers(&pokerGame.Deck, playerNum)
	pokerGame.Table = cardgame.InitTable(&pokerGame.Deck, 3)
	pokerGame.Fee = 0
	pokerGame.FirstFee = 100
	pokerGame.LastRaisedPlayerIdx = -1
	pokerGame.GameHistory = []PlayAction{}
	pokerGame.GameLogInfos = []GameLogInfo{}

	return pokerGame
}

func CalcPokerGameScore(pokerGame *PokerGame) {
	for pidx := 0; pidx < pokerGame.PlayerNum; pidx++ {
		score := GetHandScore(append(pokerGame.Table.Cards, pokerGame.Players[pidx].Hand.Cards...))
		pokerGame.Players[pidx].Score = score
		handsetStr := fmt.Sprintf("%v", EncodeHandset(score))
		pokerGame.Players[pidx].Info = handsetStr
	}
	nextscores := CalcNextHandsetScores(append(pokerGame.Table.Cards, pokerGame.Players[0].Hand.Cards...), 1)
	pokerGame.GameInfo = fmt.Sprintf("%v", nextscores)
}

func getBitidxAry(ary []int, count int) []int {
	res := make([]int, count)
	ridx := 0
	for idx := 0; idx < len(ary); idx++ {
		if ary[idx] == 1 {
			res[ridx] = idx
			ridx += 1
		}
	}
	return res
}
func isIdxaryEnd(ary []int, count int) bool {
	flg := true
	for idx := 0; idx < count; idx++ {
		flg = flg && (ary[len(ary)-1-idx] == 1)
	}
	return flg
}
func getLastBitIdx(ary []int) int {
	for idx := 0; idx < len(ary); idx++ {
		if ary[len(ary)-1-idx] == 1 {
			return len(ary) - 1 - idx
		}
	}
	return -1
}
func getIndexSet(count int, maxValue int) [][]int {
	resset := make([][]int, 0, maxValue)
	idxary := make([]int, maxValue)
	for idx := 0; idx < count; idx++ {
		idxary[idx] = 1
	}
	for !isIdxaryEnd(idxary, count) {
		resset = append(resset, getBitidxAry(idxary, count))
		targetary := idxary
		for {
			lastidx := getLastBitIdx(targetary)
			if lastidx == len(targetary)-1 {
				targetary = targetary[:len(targetary)-1]
			} else {
				idxary[lastidx+1] = 1
				idxary[lastidx] = 0
				lastNums := len(idxary) - len(targetary)
				for idx := 0; idx < lastNums; idx++ {
					idxary[len(idxary)-1-idx] = 0
				}
				for idx := 0; idx < lastNums; idx++ {
					idxary[lastidx+1+1+idx] = 1
				}
				break
			}
		}
	}
	resset = append(resset, getBitidxAry(idxary, count))
	return resset
}
func CalcNextHandsetScores(cards []cardgame.Card, addNum int) []NextCardScoreInfo {
	restdeck := cardgame.InitDeck()
	for idx := 0; idx < len(cards); idx++ {
		restdeck = cardgame.DeleteCardFromDeck(restdeck, cards[idx])
	}

	nextCardScoreInfos := make([]NextCardScoreInfo, 0)
	indexset := getIndexSet(addNum, len(restdeck.Cards))
	for idx := 0; idx < len(indexset); idx++ {
		nextcards := make([]cardgame.Card, len(cards))
		copy(nextcards, cards)
		for iidx := 0; iidx < len(indexset[idx]); iidx++ {
			card := restdeck.Cards[indexset[idx][iidx]]
			nextcards = append(nextcards, card)
		}
		score := GetHandScore(nextcards)
		nextCardScoreInfos = append(nextCardScoreInfos, NextCardScoreInfo{nextcards[len(nextcards)-addNum:], score})
	}

	return nextCardScoreInfos
}
func FilterNextHandsetScores(nscinfos []NextCardScoreInfo, yaku string) []NextCardScoreInfo {
	res := make([]NextCardScoreInfo, 0)
	handsetMap := getHandSetScoreMap()
	for idx := 0; idx < len(nscinfos); idx++ {
		if nscinfos[idx].Score > handsetMap[yaku] {
			res = append(res, nscinfos[idx])
		}
	}
	return res
}
func CalcNextHandsetProb(nscinfos []NextCardScoreInfo, yaku string) float64 {
	filteredInfo := FilterNextHandsetScores(nscinfos, yaku)
	return float64(len(filteredInfo)) / float64(len(nscinfos))
}
