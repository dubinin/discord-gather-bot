package command

type CmdInfo struct {
	*Command
}

func (c *CmdInfo) Execute() {
	game := *c.game
	if game != nil {
		if game.PlayerCount() > 0 {
			c.sendMsg("Game %s, players: %s", game.GetName(), game.GetPlayersInfo(true,
				game.PlayerLimit()-game.PlayerCount()))
		} else {
			c.sendMsg("There is no game %s", game.GetName())
		}
	}
}
