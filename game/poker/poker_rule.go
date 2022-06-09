package poker

import (
	"fmt"
	"game/cardgame"
	"math"
	"sort"
)

func getDupNums(cards []cardgame.Card, dupNum int) []int {
	cardNums := make([]int, 0, 13)
	cardNumMap := make(map[int]int)

	for cidx := 0; cidx < len(cards); cidx++ {
		cardNumMap[cards[cidx].Number]++
	}
	for key := range cardNumMap {
		if cardNumMap[key] == dupNum {
			cardNums = append(cardNums, key)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardNums)))
	return cardNums
}
func getDupSuites(cards []cardgame.Card, dupNum int) []int {
	cardNums := make([]int, 0, 5)
	cardNumMap := make(map[int]int)

	for cidx := 0; cidx < len(cards); cidx++ {
		cardNumMap[cards[cidx].Suite]++
	}
	for key := range cardNumMap {
		if cardNumMap[key] == dupNum {
			dupNums := make([]int, 0, 13*4)
			for cidx := 0; cidx < len(cards); cidx++ {
				if cards[cidx].Suite == key {
					dupNums = append(dupNums, cards[cidx].Number)
				}
			}
			sort.Sort(sort.Reverse(sort.IntSlice(dupNums)))
			cardNums = append(cardNums, dupNums[0])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardNums)))
	return cardNums
}
func getContinuousNums(cards []cardgame.Card, conNum int) []int {
	cardNums := make([]int, 0, 13)
	cardNumMap := make(map[int]int)

	for cidx := 0; cidx < len(cards); cidx++ {
		cardNumMap[cards[cidx].Number]++
	}
	for key := range cardNumMap {
		flg := true
		for idx := 1; idx < conNum; idx++ {
			_, ok := cardNumMap[key+idx]
			flg = flg && ok
		}
		if flg {
			cardNums = append(cardNums, key)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cardNums)))
	return cardNums
}
func getKabuNums(cards []cardgame.Card) []int {
	nums := make([]int, 0, 13)
	for idx := 0; idx < len(cards); idx++ {
		nums = append(nums, cards[idx].Number)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}

func IsPair(cards []cardgame.Card) int {
	cardNums := getDupNums(cards, 2)
	if len(cardNums) == 1 {
		return cardNums[0]
	}
	return -1
}
func Is2Pairs(cards []cardgame.Card) int {
	cardNums := getDupNums(cards, 2)
	if len(cardNums) == 2 {
		return cardNums[0]
	}
	return -1
}
func Is3Cards(cards []cardgame.Card) int {
	cardNums := getDupNums(cards, 3)
	if len(cardNums) == 1 {
		return cardNums[0]
	}
	return -1
}
func Is4Cards(cards []cardgame.Card) int {
	cardNums := getDupNums(cards, 4)
	if len(cardNums) == 1 {
		return cardNums[0]
	}
	return -1
}
func IsFullHouse(cards []cardgame.Card) int {
	cardPairNums := getDupNums(cards, 2)
	card3CardsNums := getDupNums(cards, 3)
	if len(cardPairNums) == 1 && len(card3CardsNums) == 1 {
		return card3CardsNums[0]
	}
	return -1
}
func IsFlush(cards []cardgame.Card) int {
	cardNums := getDupSuites(cards, 5)
	if len(cardNums) == 1 {
		return cardNums[0]
	}
	return -1
}
func IsStraight(cards []cardgame.Card) int {
	cardNums := getContinuousNums(cards, 5)
	if len(cardNums) == 1 {
		return cardNums[0]
	}
	return -1
}
func IsKabu(cards []cardgame.Card) int {
	cardNums := getKabuNums(cards)
	return cardNums[0]
}

func getHandSetMap() map[string]func([]cardgame.Card) int {
	handsets := map[string]func([]cardgame.Card) int{
		"kabu":      IsKabu,
		"pair":      IsPair,
		"2pairs":    Is2Pairs,
		"3cards":    Is3Cards,
		"straight":  IsStraight,
		"flush":     IsFlush,
		"4cards":    Is4Cards,
		"fullhouse": IsFullHouse,
	}
	return handsets
}
func getHandSetNames() []string {
	handsets := []string{
		"fullhouse",
		"4cards",
		"flush",
		"straight",
		"3cards",
		"2pairs",
		"pair",
		"kabu",
	}
	return handsets
}
func getHandSetScoreMap() map[string]int {
	BaseValue := 15
	unitScoreMap := make(map[string]int)

	handsets := getHandSetNames()
	for idx := 0; idx < len(handsets); idx++ {
		handset := handsets[idx]
		unitScoreMap[handset] = int(math.Pow(float64(BaseValue), float64(len(handsets)-idx)))
	}
	return unitScoreMap
}
func GetHandScore(cards []cardgame.Card) int {
	handsetMap := getHandSetMap()
	scoreMap := getHandSetScoreMap()
	handsets := getHandSetNames()

	score := 0
	for idx := 0; idx < len(handsets); idx++ {
		handset := handsets[idx]
		score += (handsetMap[handset](cards) + 1) * scoreMap[handset]
	}
	return score
}
func EncodeHandset(score int) []HandsetInfo {
	handsetInfo := make([]HandsetInfo, 0, len(getHandSetNames()))
	scoreMap := getHandSetScoreMap()
	handsets := getHandSetNames()

	for idx := 0; idx < len(handsets); idx++ {
		handset := handsets[idx]
		unitscore := scoreMap[handset]
		if score >= unitscore {
			handsetInfo = append(handsetInfo, HandsetInfo{handset, int(score/unitscore - 1)})
			score = score % unitscore
		}
	}

	return handsetInfo
}
func CheckAllPlayerHandset(pokerGame PokerGame) {
	table := pokerGame.Table
	players := pokerGame.Players

	for pidx := 0; pidx < len(players); pidx++ {
		score := GetHandScore(append(table.Cards, players[pidx].Hand.Cards...))
		fmt.Println(EncodeHandset(score))
	}
}
