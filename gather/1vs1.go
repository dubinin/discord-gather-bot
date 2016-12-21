package gather

type OneVsOne struct {
	*Game
}

func (g *OneVsOne) GetName() string {
	return "1v1"
}

func (g *OneVsOne) PlayerLimit() int {
	return 2
}

func (g *OneVsOne) Complete() string {
	return g.GetShuffle(g.GetName(), g.PlayerLimit())
}
