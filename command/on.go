package command

type CmdOn struct {
	*Command
}

func (c *CmdOn) Execute() {
	game := *c.game
	if game != nil {
		if game.AddPlayer(c.author, game.PlayerLimit()) {
			c.sendMsg("<@%s> create game %s, players %d/%d",
				c.author.User.ID, game.GetName(), game.PlayerCount(), game.PlayerLimit())
			c.updateStatus()
		} else {
			c.sendMsg("<@%s> You already in the game", c.author.User.ID)
		}
	}
}
