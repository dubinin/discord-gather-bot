package gather

import "fmt"

type ThreeS struct {
	*Game
}

func (g *ThreeS) GetName() string {
	return "3s"
}

func (g *ThreeS) PlayerLimit() int {
	return 3
}

func (g *ThreeS) Complete() string {
	return fmt.Sprintf("Team for ranked game %s complete %s",
		g.GetName(), g.GetPlayersInfo(false, 0))
}
