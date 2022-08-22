package structs

import (
	"github.com/bwmarrin/discordgo"
)

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

var Commands []Command

// Creates a new command object.
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

	Commands = append(Commands, c)

	return &c
}

// Checks if a command exists.
// TODO: Make this work with aliases.
func CommandExists(name string) bool {
	for _, c := range Commands {
		if c.Name == name {
			return true
		}
	}
	return false
}

// TODO: Make this work with aliases.
func GetCommand(name string) *Command {
	for _, c := range Commands {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

func fallback(v, f string) string {
	if len(v) == 0 {
		return f
	}
	return v
}
