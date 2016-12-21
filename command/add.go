package command

type CmdAdd struct {
	*Command
}

func (c *CmdAdd) Execute() {
	game := *c.game
	if game != nil {
		if game.AddPlayer(c.author, game.PlayerLimit()) {
			freeSlots := game.PlayerLimit() - game.PlayerCount()
			if freeSlots > 0 {
				c.sendMsg("<@%s> enter into game %s, free slots %d", c.author.User.ID,
					game.GetName(), freeSlots)
				c.updateStatus()
			} else {
				c.sendMsg(game.Complete())
				c.gather.CloseGame(game)
				c.updateStatus()
			}
		} else {
			c.sendMsg("<@%s> already in the game %s", c.author.User.ID, game.GetName())
		}
	}
}
