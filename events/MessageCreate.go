package events

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by bots. This also makes the bot ignore itself.
	if m.Author.Bot {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "go.ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong! "+s.HeartbeatLatency().String())
	}

	if m.Content == "go.invite" {
		s.ChannelMessageSend(m.ChannelID, "https://discord.com/oauth2/authorize?scope=applications.commands%20bot&permissions=268561488&client_id=1007010887447625748")
	}

	if m.Content == "go.help" {
		s.ChannelMessageSend(m.ChannelID, "```go\nCommands:\ngo.ping\ngo.invite\ngo.help\n```")
	}
}
