package poker

type PlayAction struct {
	PlayerID   int
	ActionType string
	Cost       int
	Phase      int
}

func NPCPlay(pokerGame PokerGame, playerId int) PlayAction {
	action := PlayAction{}
	action.PlayerID = playerId
	action.ActionType = "hold"
	action.Phase = pokerGame.Phase

	player := pokerGame.Players[playerId]
	strategy := player.Strategy
	if strategy == "strategy_0" {
		action = StrategyCheckerBet(playerId, pokerGame)
	}
	if strategy == "strategy_1" {
		action = StrategyCurrentAboveHandTimesBet(playerId, pokerGame, "pair", 2)
	}
	if strategy == "strategy_2" {
		action = StrategyCurrentAboveHandTimesBet(playerId, pokerGame, "2pairs", 5)
	}
	if strategy == "strategy_3" {
		action = StrategyCurrentAboveHandOffsetBet(playerId, pokerGame, "2pairs", 2000)
	}
	if strategy == "strategy_4" {
		action = StrategyNextAboveHandTimesBet(playerId, pokerGame, "pair", 2)
	}
	if strategy == "strategy_5" {
		action = StrategyNextAboveHandTimesBet(playerId, pokerGame, "2pairs", 5)
	}
	if strategy == "strategy_6" {
		action = StrategyNextAboveHandOffsetBet(playerId, pokerGame, "2pairs", 2000)
	}
	if strategy == "strategy_7" {
		action = StrategyCheckerBet(playerId, pokerGame)
	}
	if strategy == "strategy_8" {
		action = StrategyCheckerBet(playerId, pokerGame)
	}
	if strategy == "strategy_9" {
		action = StrategyCheckerBet(playerId, pokerGame)
	}
	if strategy == "strategy_raiser" {
		action = StrategyRaiserBet(playerId, pokerGame)
	}
	if strategy == "strategy_checker" {
		action = StrategyCheckerBet(playerId, pokerGame)
	}
	return action
}

func StrategyCurrentAboveHandTimesBet(player_Id int, pokerGame PokerGame, hand_name string, times int) PlayAction {
	fee := pokerGame.Fee
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	bet := fee
	cur_score := GetHandScore(append(pokerGame.Table.Cards, player.Hand.Cards...))
	scoreMap := getHandSetScoreMap()
	cri_score := scoreMap[hand_name]
	is_raise := cur_score >= cri_score
	if is_raise {
		bet = fee * times
		if fee == 0 {
			bet = 500
		}
		action.ActionType = "raise"
		action.Cost = bet
		if player.Fund-bet <= 0 {
			action.ActionType = "all_in"
			action.Cost = player.Fund
		}
	}
	if !is_raise && fee >= 1000 {
		action.ActionType = "hold"
	}

	return action
}
func StrategyCurrentAboveHandOffsetBet(player_Id int, pokerGame PokerGame, hand_name string, offset int) PlayAction {
	fee := pokerGame.Fee
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	bet := fee
	cur_score := GetHandScore(append(pokerGame.Table.Cards, player.Hand.Cards...))
	scoreMap := getHandSetScoreMap()
	cri_score := scoreMap[hand_name]
	is_raise := cur_score >= cri_score
	if is_raise {
		bet = fee + offset
		if fee == 0 {
			bet = 500
		}
		action.ActionType = "raise"
		action.Cost = bet
		if player.Fund-bet <= 0 {
			action.ActionType = "all_in"
			action.Cost = player.Fund
		}
	}
	if !is_raise && fee >= 1000 {
		action.ActionType = "hold"
	}

	return action
}

func StrategyNextAboveHandTimesBet(player_Id int, pokerGame PokerGame, hand_name string, times int) PlayAction {
	fee := pokerGame.Fee
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	bet := fee
	nextscores := CalcNextHandsetScores(append(pokerGame.Table.Cards, player.Hand.Cards...), 1)
	pair_prob := CalcNextHandsetProb(nextscores, hand_name)
	is_raise := pair_prob >= 0.5
	if is_raise {
		bet = fee * times
		if fee == 0 {
			bet = 500
		}
		action.ActionType = "raise"
		action.Cost = bet
		if player.Fund-bet <= 0 {
			action.ActionType = "all_in"
			action.Cost = player.Fund
		}
	}
	if !is_raise && fee >= 1000 {
		action.ActionType = "hold"
	}

	return action
}
func StrategyNextAboveHandOffsetBet(player_Id int, pokerGame PokerGame, hand_name string, offset int) PlayAction {
	fee := pokerGame.Fee
	player := pokerGame.Players[player_Id]
	action := PlayAction{}
	action.PlayerID = player_Id
	action.ActionType = "check"
	action.Cost = fee
	action.Phase = pokerGame.Phase

	bet := fee
	nextscores := CalcNextHandsetScores(append(pokerGame.Table.Cards, player.Hand.Cards...), 1)
	pair_prob := CalcNextHandsetProb(nextscores, hand_name)
	is_raise := pair_prob >= 0.5
	if is_raise {
		bet = fee + offset
		if fee == 0 {
			bet = 500
		}
		action.ActionType = "raise"
		action.Cost = bet
		if player.Fund-bet <= 0 {
			action.ActionType = "all_in"
			action.Cost = player.Fund
		}
	}
	if !is_raise && fee >= 1000 {
		action.ActionType = "hold"
	}

	return action
}

func StrategyRaiserBet(player_Id int, pokerGame PokerGame) PlayAction {
	fee := pokerGame.Fee
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

func StrategyCheckerBet(player_Id int, pokerGame PokerGame) PlayAction {
	fee := pokerGame.Fee
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
