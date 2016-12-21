package command

type CmdHelp struct {
	*Command
}

func (c *CmdHelp) Execute() {
	c.sendMsg("<@%s> try execute command !help", c.author.User.ID)
}
