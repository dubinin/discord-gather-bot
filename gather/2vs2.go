package gather

type TwoVsTwo struct {
	*Game
}

func (g *TwoVsTwo) GetName() string {
	return "2v2"
}

func (g *TwoVsTwo) PlayerLimit() int {
	return 4
}

func (g *TwoVsTwo) Complete() string {
	return g.GetShuffle(g.GetName(), g.PlayerLimit())
}
