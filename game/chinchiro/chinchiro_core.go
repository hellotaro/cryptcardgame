package chinchiro

import (
	"game/dicegame"
)

type ChinchiroGame struct {
	Phase          int
	ChinchiroDices dicegame.DiceSet
}

func InitChinchiroGame() ChinchiroGame {
	game := ChinchiroGame{0, dicegame.GetNewDiceSet(3)}
	return game
}
