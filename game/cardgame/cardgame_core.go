package cardgame

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Card struct {
	Suite  int
	Number int
}

func (c Card) getString() string {
	var suite_str string
	switch c.Suite {
	case 0:
		suite_str = "S"
	case 1:
		suite_str = "C"
	case 2:
		suite_str = "D"
	case 3:
		suite_str = "H"
	}
	return "(" + suite_str + "," + strconv.Itoa(c.Number) + ")"
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Cards []Card
}

type Table struct {
	Cards []Card
	Pod   int
}

type Player struct {
	Name     string
	Hand     Hand
	Strategy string
	State    string
	Score    int
	Rank     int
	Fund     int
	Bet      int
	Info     string
}
type Players []Player

func (p Players) Len() int {
	return len(p)
}
func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Players) Less(i, j int) bool {
	if p[i].State == "dead" && p[j].State == "live" {
		return true
	}
	if p[i].State == "live" && p[j].State == "dead" {
		return false
	}
	return p[i].Score < p[j].Score
}

func GetCardsStr(cards []Card) string {
	resstr := ""
	for idx := 0; idx < len(cards); idx++ {
		resstr += cards[idx].getString() + " "
	}
	return resstr
}

func InitDeck() Deck {
	deck := Deck{}
	deck.Cards = make([]Card, 13*4)
	for suite := 0; suite < 4; suite++ {
		for number := 0; number < 13; number++ {
			deck.Cards[suite*13+number] = Card{suite, number}
		}
	}
	return deck
}
func DrawCardFromDeck(deck *Deck) Card {
	rand.Seed(time.Now().UnixNano())
	card_idx := rand.Intn(len(deck.Cards))
	card := deck.Cards[card_idx]
	deck.Cards = append(deck.Cards[:card_idx], deck.Cards[card_idx+1:]...)
	return card
}
func DeleteCardFromDeck(deck Deck, card Card) Deck {
	for idx := 0; idx < len(deck.Cards); idx++ {
		if deck.Cards[idx].Suite == card.Suite && deck.Cards[idx].Number == card.Number {
			deck.Cards = append(deck.Cards[:idx], deck.Cards[idx+1:]...)
		}
	}
	return deck
}
func PrintDeck(deck Deck) {
	fmt.Println("===Deck===")
	prev_suite := -1
	for idx := 0; idx < len(deck.Cards); idx++ {
		card := deck.Cards[idx]
		if prev_suite != card.Suite {
			if prev_suite != -1 {
				fmt.Printf("\n")
			}
			prev_suite = card.Suite
		}
		fmt.Printf("%v ", card.getString())
	}
	fmt.Printf("\n")
}

func InitTable(deck *Deck, tableNum int) Table {
	table := Table{}
	table.Cards = make([]Card, tableNum, 5)
	for idx := 0; idx < tableNum; idx++ {
		table.Cards[idx] = DrawCardFromDeck(deck)
	}
	table.Pod = 0
	return table
}
func AddCardToTable(table *Table, deck *Deck) {
	table.Cards = append(table.Cards, DrawCardFromDeck(deck))
}
func PrintTable(table Table) {
	fmt.Println("===Table===")
	cards := table.Cards
	for idx := 0; idx < len(cards); idx++ {
		fmt.Printf("%v ", cards[idx].getString())
	}
	fmt.Printf("\n")
	fmt.Printf("[Pod] %v\n", table.Pod)
}

func PrintPlayer(player Player, table Table) {
	fmt.Printf("> %v (%v)\n", player.Name, player.State)
	fmt.Printf("[Fund(Bet)] %v(%v)\n", (player.Fund - player.Bet), player.Bet)
	fmt.Printf("[Hand] ")
	for cidx := 0; cidx < len(player.Hand.Cards); cidx++ {
		fmt.Printf("%v ", player.Hand.Cards[cidx].getString())
	}
	fmt.Printf("\n")
}
func PrintPlayersBet(players []Player) {
	fmt.Println(">Current Bet")
	for pidx := 0; pidx < len(players); pidx++ {
		fmt.Printf("%v - Fund: %v, Bet: %v\n", players[pidx].Name, (players[pidx].Fund - players[pidx].Bet), players[pidx].Bet)
	}
}

func GetRankedPlayers(players Players) []Player {
	sort.Sort(sort.Reverse(players))
	return players
}

func SetPlayerRankingIndex(players Players) []Player {
	rankedPlayers := GetRankedPlayers(players)
	for idx := 0; idx < len(rankedPlayers); idx++ {
		for pidx := 0; pidx < len(players); pidx++ {
			if players[pidx].Name == rankedPlayers[idx].Name {
				players[pidx].Rank = idx + 1
			}
		}
	}
	return players
}
