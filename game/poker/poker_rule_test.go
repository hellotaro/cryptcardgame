package poker

import (
	"game/cardgame"
	"testing"
)

func TestPokerRules(t *testing.T) {
	var cards []cardgame.Card

	cards = make([]cardgame.Card, 0, 13*4)
	cards = append(cards, cardgame.Card{0, 0})
	cards = append(cards, cardgame.Card{0, 1})
	cards = append(cards, cardgame.Card{0, 2})
	cards = append(cards, cardgame.Card{0, 3})
	cards = append(cards, cardgame.Card{0, 4})

	if IsStraight(cards) != 0 {
		t.Fatal("judge failed!: straight")
	}

	cards[1] = cardgame.Card{1, 0}
	if IsPair(cards) != 0 {
		t.Fatal("judge failed!: pair")
	}

	cards[2] = cardgame.Card{2, 0}
	if Is3Cards(cards) != 0 {
		t.Fatal("judge failed!: 3cards")
	}

	cards[3] = cardgame.Card{3, 0}
	if Is4Cards(cards) != 0 {
		t.Fatal("judge failed!: 4cards")
	}

	cards[3] = cardgame.Card{0, 1}
	cards[4] = cardgame.Card{1, 1}
	if IsFullHouse(cards) != 0 {
		t.Fatal("judge failed!: FullHouse")
	}

	cards = make([]cardgame.Card, 0, 13*4)
	cards = append(cards, cardgame.Card{0, 0})
	cards = append(cards, cardgame.Card{0, 1})
	cards = append(cards, cardgame.Card{0, 2})
	cards = append(cards, cardgame.Card{0, 3})
	cards = append(cards, cardgame.Card{0, 4})
	score := GetHandScore(cards)
	handsetInfo := EncodeHandset(score)
	flg := false
	for idx := 0; idx < len(handsetInfo); idx++ {
		flg = flg || handsetInfo[idx].Yaku == "straight"
	}
	if !flg {
		t.Fatal("score failed!: straight")
	}
	flg = false
	for idx := 0; idx < len(handsetInfo); idx++ {
		flg = flg || handsetInfo[idx].Yaku == "flush"
	}
	if !flg {
		t.Fatal("score failed!: flush")
	}
}
