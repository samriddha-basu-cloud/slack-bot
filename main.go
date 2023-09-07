package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}

}

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5887005301840-5864184421762-8LV02q6YfmXujFhe8DouNzlp")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05RDV6KRM1-5876733714609-f3d0b6bbcadfc8fb165e84b70e8587e47fe1b594beb15f3d962a446ca49d7b12")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("Hi", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("নমস্কার")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
