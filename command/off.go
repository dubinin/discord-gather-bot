package command

type CmdOff struct {
	*Command
}

func (c *CmdOff) Execute() {
	game := *c.game
	if game != nil {
		if c.gather.IsAdmin(c.author.User.ID) {
			c.gather.CloseGame(game)
			c.sendMsg("<@%s> close game %s", c.author.User.ID, game.GetName())
			c.updateStatus()
		}
	}
}
