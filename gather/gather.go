package gather

import "time"

type Gather struct {
	Games    map[string]Gatherable
	ChanelID string
	Admins   []string
}

func (g *Gather) CloseGame(game Gatherable) {
	delete(g.Games, game.GetName())
}

func (g *Gather) AddGame(game Gatherable) {
	if game != nil {
		g.Games[game.GetName()] = game
	}
}

func (g *Gather) GetGame(gameName string) *Gatherable {
	game := g.Games[gameName]
	if game != nil {
		return &game
	}
	return nil
}

func (g *Gather) CreateGame(gameName string, initiator Player) *Gatherable {
	var newGame Gatherable
	if len(gameName) > 0 {
		game := &Game{
			Start:     time.Now(),
			Initiator: initiator,
		}
		switch gameName {
		case "1v1":
			newGame = &OneVsOne{game}
		case "2v2":
			newGame = &TwoVsTwo{game}
		case "3v3":
			newGame = &ThreeVsThree{game}
		case "3s":
			newGame = &ThreeS{game}
		case "2s":
			newGame = &TwoS{game}
		}
	}
	g.AddGame(newGame)
	return &newGame
}

func (g *Gather) IsAdmin(userID string) bool {
	if len(g.Admins) == 0 {
		return true
	}
	for _, admin := range g.Admins {
		if admin == userID {
			return true
		}
	}
	return false
}
