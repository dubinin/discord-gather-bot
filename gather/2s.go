package gather

import "fmt"

type TwoS struct {
	*Game
}

func (g *TwoS) GetName() string {
	return "2s"
}

func (g *TwoS) PlayerLimit() int {
	return 2
}

func (g *TwoS) Complete() string {
	return fmt.Sprintf("Team for ranked game %s complete %s",
		g.GetName(), g.GetPlayersInfo(false, 0))
}
