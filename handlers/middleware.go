package handlers

import (
	"context"
	"log"
	// "os"
	// "os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)


func ShowMessageWithUserID(next bot.HandlerFunc) bot.HandlerFunc  {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("%d say: %s", update.Message.From.ID, update.Message.Text)
		}
		next(ctx, b, update)
	}
}


func ShowMessageWithUserName(next bot.HandlerFunc) bot.HandlerFunc  {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("%s say: %s", update.Message.From.Username, update.Message.Text)
		}

		next(ctx, b, update)
	}
}