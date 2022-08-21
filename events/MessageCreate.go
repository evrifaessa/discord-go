package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/evrifaessa/discord-go/commands"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by bots. This also makes the bot ignore itself.
	if m.Author.Bot {
		return
	}

	// TODO: Register the commands to a map and dynamically load them.
	if m.Content == "go.ping" {
		commands.Ping.Run(s, m)
	}

	if m.Content == "go.invite" {
		commands.Invite.Run(s, m)
	}

	if m.Content == "go.help" {
		commands.Help.Run(s, m)
	}
}
