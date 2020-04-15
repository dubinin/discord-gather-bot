package command

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dubinin/discord-gather-bot/gather"
	"github.com/bwmarrin/discordgo"

	log "github.com/Sirupsen/logrus"
)

type Executable interface {
	Execute()
}

type Command struct {
	content   string
	author    gather.Player
	channelID string
	game      *gather.Gatherable
	session   *discordgo.Session
	gather    *gather.Gather
}

func New(session *discordgo.Session, msg *discordgo.MessageCreate,
	globalGather *gather.Gather) Executable {
	var op, content = parseCommand(strings.ToLower(msg.Content))
	cmd := &Command{
		content:   content,
		author:    gather.Player{User: *msg.Author},
		channelID: msg.ChannelID,
		session:   session,
		gather:    globalGather,
	}
	log.Debugf("Parsed command, operator [%v], content [%v]", op, content)
	game := globalGather.GetGame(content)
	if game != nil {
		cmd.game = game
	} else {
		cmd.game = globalGather.CreateGame(content, cmd.author)
	}
	switch op {
	case "!on":
		return &CmdOn{cmd}
	case "!off":
		return &CmdOff{cmd}
	case "!add":
		return &CmdAdd{cmd}
	case "!del":
		return &CmdDel{cmd}
	case "!info":
		return &CmdInfo{cmd}
	case "!help":
		return &CmdHelp{cmd}
	}
	return cmd
}

func (c Command) String() string {
	return fmt.Sprintf("Cmd: %s", c.content)
}

func (c *Command) Execute() {}

func (c *Command) sendMsg(format string, a ...interface{}) {
	channel := c.channelID
	if len(c.gather.ChanelID) > 0 {
		channel = c.gather.ChanelID
	}
	c.session.ChannelMessageSend(channel, fmt.Sprintf(format, a...))
}

func (c *Command) updateStatus() {
	var buffer bytes.Buffer
	for gameName := range c.gather.Games {
		if gameName != "game" {
			buffer.WriteString(fmt.Sprintf("%s ", gameName))
		}
	}
	c.session.UpdateStatus(0, buffer.String())
}

func parseCommand(msg string) (string, string) {
	splitMsg := strings.Split(msg, " ")
	splitLen := len(splitMsg)
	switch {
	case splitLen == 1:
		return splitMsg[0], ""
	case splitLen >= 2:
		return splitMsg[0], splitMsg[1]
	default:
		return "", ""
	}
}
