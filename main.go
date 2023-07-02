package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"lido_dao_bot/handlers"

	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/go-telegram/bot"
)

func main(){
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(handlers.ShowMessageWithUserName),
		bot.WithDefaultHandler(handlers.ProcessMessage),
		// bot.WithMessageTextHandler("text", bot.MatchTypePrefix, handlers.ProcessMessage),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, handlers.ProcessCallbackQueryMessage),
	}

	apiBot, err := bot.New("5885818251:AAHTAP7gBY8caUNRriJByHk7UQbwicVW9E0", opts...)
	if nil != err {
		log.Panic(err)
	}
	apiBot.Start(ctx)
}