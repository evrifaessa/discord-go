package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/evrifaessa/discord-go/structs"
)

var Prefix (string) = "go."

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by bots. This also makes the bot ignore itself.
	if m.Author.Bot {
		return
	}

	// TODO: Register the commands to a map and dynamically load them.
	if m.Content == Prefix+"ping" {
		structs.GetCommand("ping").Run(s, m)
	}

	if m.Content == Prefix+"invite" {
		structs.GetCommand("invite").Run(s, m)
	}

	if m.Content == Prefix+"help" {
		structs.GetCommand("help").Run(s, m)
	}
}
