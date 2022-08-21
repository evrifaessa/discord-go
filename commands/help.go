package commands

import (
	"github.com/bwmarrin/discordgo"
)

var HelpConfig = map[string]string{
	"name": "help",
}

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "```go\nCommands:\ngo.ping\ngo.invite\ngo.help\n```")
}
