package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/evrifaessa/discord-go/events"
	"github.com/evrifaessa/discord-go/structs"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", LookupEnvOrString("TOKEN", ""), "Bot Token")
	flag.Parse()
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	fmt.Println("Successfully loaded " + strconv.Itoa(len(structs.Commands)) + " commands.")

	/* Let's load our events into the memory.
	 * Our events are in the events/ folder. Each file in this folder is a separate event.
	 * Check the DiscordGo documentation for more information on how the events are named.
	 * https://pkg.go.dev/github.com/bwmarrin/discordgo#Session.AddHandler
	 */

	// TODO: Dynamically load events and not hardcode them.
	dg.AddHandler(events.MessageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
