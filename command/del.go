package command

type CmdDel struct {
	*Command
}

func (c *CmdDel) Execute() {
	game := *c.game
	if game != nil {
		if game.RemovePlayer(c.author) {
			freeSlots := game.PlayerLimit() - game.PlayerCount()
			if freeSlots == game.PlayerLimit() {
				c.sendMsg("Close game %s", game.GetName())
				c.gather.CloseGame(game)
				c.updateStatus()
			} else {
				c.sendMsg("<@%s> leave game %s, free slots %d",
					c.author.User.ID, game.GetName(), freeSlots)
			}
		}
	}
}
