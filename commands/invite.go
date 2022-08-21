package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/evrifaessa/discord-go/structs"
)

var Invite = structs.NewCommand(
	"invite",   // Name
	"misc",     // Category
	[]string{}, // Aliases
	true,       // Enabled
	false,      // ArgsRequired
	false,      // Hidden

	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "https://discord.com/oauth2/authorize?scope=applications.commands%20bot&permissions=268561488&client_id=1007010887447625748")
	},
)
