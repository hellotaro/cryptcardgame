package poker

type PlayAction struct {
	PlayerID   int
	ActionType string
	Cost       int
	Phase      int
}

func NPCPlay(pokerGame PokerGame, playerId int, fee int) PlayAction {
	action := PlayAction{}
	action.PlayerID = playerId
	action.ActionType = "hold"
	action.Phase = pokerGame.Phase

	player := pokerGame.Players[playerId]
	strategy := player.Strategy
	if strategy == "strategy_0" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_1" {
		action = StrategyAbovePairBet(playerId, fee, pokerGame)
		//action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_2" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_3" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_4" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_5" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_6" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_7" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_8" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_9" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_raiser" {
		action = StrategyRaiserBet(playerId, fee, pokerGame)
	}
	if strategy == "strategy_checker" {
		action = StrategyCheckerBet(playerId, fee, pokerGame)
	}
	return action
}

func StrategyAbovePairBet(player_Id int, fee int, pokerGame PokerGame) PlayAction {
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	bet := fee
	nextscores := CalcNextHandsetScores(append(pokerGame.Table.Cards, player.Hand.Cards...), 1)
	pair_prob := CalcNextHandsetProb(nextscores, "pair")
	if pair_prob >= 0.5 {
		bet = fee * 2
		if fee == 0 {
			bet = 100
		}
		action.ActionType = "raise"
		action.Cost = bet
	}
	if player.Fund-bet <= 0 {
		action.ActionType = "all_in"
		action.Cost = player.Fund
	}

	return action
}

func StrategyRaiserBet(player_Id int, fee int, pokerGame PokerGame) PlayAction {
	player := pokerGame.Players[player_Id]

	offset := 500

	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "raise"
	action.Cost = fee + offset
	action.Phase = pokerGame.Phase

	if player.Fund-(fee+offset) <= 0 {
		action.ActionType = "check"
		action.Cost = fee
	}
	if player.Fund-player.Bet-fee <= 0 {
		action.ActionType = "all_in"
		action.Cost = player.Fund
	}
	return action
}

func StrategyCheckerBet(player_Id int, fee int, pokerGame PokerGame) PlayAction {
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	if player.Fund-fee <= 0 {
		action.ActionType = "all_in"
		action.Cost = player.Fund
	}

	return action
}
