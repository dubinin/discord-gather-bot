package gather

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Gatherable interface {
	Complete() string
	AddPlayer(player Player, limit int) bool
	RemovePlayer(player Player) bool
	PlayerLimit() int
	PlayerCount() int
	GetName() string
	GetPlayersInfo(silent bool, emptySlot int) string
}

type Game struct {
	Start     time.Time
	Initiator Player
	Players   []Player
}

func (g *Game) AddPlayer(player Player, limit int) bool {
	if limit > g.PlayerCount() {
		if !g.isPlayerIn(player) {
			g.Players = append(g.Players, player)
			return true
		}
	}
	return false
}

func (g *Game) RemovePlayer(player Player) bool {
	for i, existed := range g.Players {
		if existed.User.ID == player.User.ID {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return !g.isPlayerIn(player)
		}
	}
	return false
}

func (g *Game) isPlayerIn(player Player) bool {
	for _, existed := range g.Players {
		if existed.User.ID == player.User.ID {
			return true
		}
	}
	return false
}

func (g *Game) PlayerLimit() int {
	return 0
}

func (g *Game) PlayerCount() int {
	return len(g.Players)
}

func (g *Game) Complete() string { return "" }

func (g *Game) GetName() string {
	return "game"
}

func (g *Game) GetShuffle(gameName string, limit int) string {
	shuffle := make([]Player, limit)
	perm := rand.Perm(limit)
	for i, v := range perm {
		shuffle[v] = g.Players[i]
	}
	middle := limit / 2
	orange := GetPlayerList(shuffle[:middle], 0, false)
	blue := GetPlayerList(shuffle[middle:], 0, false)
	pw := RandSeq(3)
	return fmt.Sprintf("Orange %s vs Blue %s, server: **%s-%s**, pw: **%s**, <@%s> will make server",
		orange, blue, gameName, pw, pw, shuffle[0].User.ID)
}

func (g *Game) GetPlayersInfo(silent bool, emptySlot int) string {
	return GetPlayerList(g.Players, emptySlot, silent)
}

func GetPlayerList(players []Player, emptySlots int, silent bool) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for _, p := range players {
		if silent {
			buffer.WriteString(fmt.Sprintf(" **%s** ", p.User.Username))
		} else {
			buffer.WriteString(fmt.Sprintf(" <@%s> ", p.User.ID))
		}
	}
	if emptySlots > 0 {
		for ; emptySlots > 0; emptySlots-- {
			buffer.WriteString(" x ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
