# evrifaessa/discord-go
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/evrifaessa/discord-go)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/evrifaessa)](https://github.com/sponsors/evrifaessa)
![Lines of code](https://img.shields.io/tokei/lines/github/evrifaessa/discord-go)

A very basic Discord bot written entirely in [Go](https://go.dev) using [DiscordGo](https://github.com/bwmarrin/discordgo). I started this project because I have no knowledge of Go and think the best way to learn something is to develop a project. PRs appreciated, I will gradually improve this project in my free time.

I also need a better name for the project.

Invite the bot [here](https://discord.com/oauth2/authorize?scope=applications.commands%20bot&permissions=268561488&client_id=1007010887447625748).

## How to run locally
You need to have Go installed. [Here](https://www.digitalocean.com/community/tutorial_collections/how-to-install-go) is a tutorial from DigitalOcean.

### Create an application if you already haven't
Create an application over the Discord Developer Portal and create a bot account associated with that application. There are plenty of guides for this on the internet.

### Run the bot using your own token
Run `go run main.go -t TOKEN` in your terminal. Of course, you need to replace "TOKEN" with the token you got from the Discord Developer Portal.
To make things much easier, you can also set your token as an environment variable named `TOKEN` and run the command without the `-t` argument.

## Questions
* [Create an issue on GitHub](https://github.com/evrifaessa/discord-go/issues/new/choose) 
* Discord: `evrifaessa#0001` (ID: 650001751050551327)
* Twitter: [yag1zhantw](https://twitter.com/yag1zhantw/)
* E-mail: [hi@yagiz.dev](mailto:hi@yagiz.dev)

## Licensing
evrifaessa/discord-go is open-source software and is released under the Apache v2.0 license. See [LICENSE](LICENSE) for the full license terms.
