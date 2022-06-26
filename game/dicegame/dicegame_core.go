package dicegame

import (
	"math/rand"
	"time"
)

type Dice struct {
	Number int
}

type DiceSet struct {
	Dices []Dice
}

func GetNewDice() Dice {
	return Dice{0}
}

func GetNewDiceSet(diceNum int) DiceSet {
	var diceSet DiceSet
	diceSet.Dices = make([]Dice, diceNum)
	return diceSet
}

func RollDiceSet(diceSet DiceSet) DiceSet {
	rand.Seed(time.Now().UnixNano())
	for idx := 0; idx < len(diceSet.Dices); idx++ {
		diceNum := rand.Intn(6) + 1
		diceSet.Dices[idx].Number = diceNum
	}
	return diceSet
}

func GetDiceSetSum(diceSet DiceSet) int {
	sum := 0
	for idx := 0; idx < len(diceSet.Dices); idx++ {
		sum += diceSet.Dices[idx].Number
	}
	return sum
}
