package gather

type ThreeVsThree struct {
	*Game
}

func (g *ThreeVsThree) GetName() string {
	return "3v3"
}

func (g *ThreeVsThree) PlayerLimit() int {
	return 6
}

func (g *ThreeVsThree) Complete() string {
	return g.GetShuffle(g.GetName(), g.PlayerLimit())
}
