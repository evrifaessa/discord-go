package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/evrifaessa/discord-go/structs"
)

var Help = structs.NewCommand(
	"help",     // Name
	"misc",     // Category
	[]string{}, // Aliases
	true,       // Enabled
	false,      // ArgsRequired
	false,      // Hidden

	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "```go\nCommands:\ngo.ping\ngo.invite\ngo.help\n```")
	},
)
