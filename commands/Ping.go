package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/evrifaessa/discord-go/structs"
)

var Ping = structs.NewCommand(
	"ping",     // Name
	"misc",     // Category
	[]string{}, // Aliases
	true,       // Enabled
	false,      // ArgsRequired
	false,      // Hidden

	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "Pong! "+s.HeartbeatLatency().String())
	},
)
