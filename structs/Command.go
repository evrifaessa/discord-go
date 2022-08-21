package structs

import "github.com/bwmarrin/discordgo"

type Runner func(s *discordgo.Session, m *discordgo.MessageCreate)

type Command struct {
	Name         string
	Aliases      []string
	Category     string
	Enabled      bool
	ArgsRequired bool
	Hidden       bool
	Run          Runner
}

func NewCommand(name, category string, aliases []string, enabled, argsRequired, hidden bool, f Runner) *Command {
	c := Command{
		Name:         name,
		Aliases:      aliases,
		Category:     fallback(category, "misc"),
		Enabled:      enabled,
		ArgsRequired: argsRequired,
		Hidden:       hidden,
		Run:          f,
	}
	return &c
}

func fallback(v, f string) string {
	if len(v) == 0 {
		return f
	}
	return v
}
