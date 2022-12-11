package main

import (
	"gengar/commands"
	"gengar/embeds"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
)

var (
	log   = logrus.New()
	token = os.Getenv("DISCORD_BOT_TOKEN")
)

func main() {
	client := disgord.New(disgord.Config{
		Intents:  disgord.AllIntents(),
		BotToken: token,
		Logger:   log,
	})
	defer func(gateway disgord.GatewayQueryBuilder) {
		if err := gateway.StayConnectedUntilInterrupted(); err != nil {
			panic(err)
		}
	}(client.Gateway())
	router := commands.Create(client)
	router.RegisterCMDList(&commands.Command{
		Name:        "testd",
		Description: "bleh",
		Handler: func(ctx *commands.Ctx) {
			err := ctx.Response(&commands.ResponseOptions{
				Flags: int(disgord.MessageFlagEphemeral),
				Embeds: []*disgord.Embed{
					embeds.CreateEmbed().SetColor("0x00ffff").SetDescription("hi").Build(),
				},
			})
			if err != nil {
				panic(err)
			}
		},
	})
	router.Init()
}
