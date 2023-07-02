package handlers

import (
	"context"
	"log"
	// "strings"

	// "os"
	// tb "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)



var supportedNetwors = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Ethereum", CallbackData: "ether"},
			{Text: "Polygon", CallbackData: "polygon"},
		}, 
		{
			{Text: "Solana", CallbackData: "solana"},
			{Text: "Polkadot", CallbackData: "polkadot"},
		},
		{
			{Text: "Kusama", CallbackData: "kusama"},
		},
	},
}

func ProcessMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch update.Message.Text {
	case "/start":
		reply := "Welcome to"
		SendTextMessage(ctx, b, update.Message.Chat.ID, reply)

	case "/help":
		reply := "Welcome to"
		SendTextMessage(ctx, b, update.Message.Chat.ID, reply)

	case "/stake":
		reply := "Lido lets you stake tokens from many networks. Choose a network below to get started."
		SendInlineKeyboard(ctx, b, update.Message.Chat.ID, reply, supportedNetwors)
	
	}
}


func ProcessCallbackQueryMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	
	switch update.CallbackQuery.Data {
	case "button_1":
		reply := "Callback Query reply to 1"
		SendCallbackQueryMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply )
	case "button_2":
		reply := "Callback Query reply to 2"
		SendCallbackQueryMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply )
	case "button_3":
		reply := "Callback Query reply to 3"
		SendCallbackQueryMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply )
	case "button_4":
		reply := "Callback Query reply to 4"
		SendCallbackQueryMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply )
	}
}


func SendInlineKeyboard(ctx context.Context, b *bot.Bot, ChatID int64, text string, kb *models.InlineKeyboardMarkup, ReplyToMessageID ...int) {

	params := &bot.SendMessageParams{
		ChatID:      ChatID,
		Text:        text,
		ReplyMarkup: kb,
	}

	if len(ReplyToMessageID) > 0 {
		params.ReplyToMessageID = ReplyToMessageID[0]
	}

	_, err := b.SendMessage(ctx, params)
	if nil != err {
		log.Panic(err)
	}
}


func SendTextMessage(ctx context.Context, b *bot.Bot, ChatID int64, text string, ReplyToMessageID ...int){
	messageParams := &bot.SendMessageParams{
		ChatID:      ChatID,
		Text:        text,
	}

	if len(ReplyToMessageID) > 0 {
		messageParams.ReplyToMessageID = ReplyToMessageID[0]
	}

	_, err := b.SendMessage(ctx, messageParams)
	if err != nil {
		log.Panic(err)
	}
}


func SendParsedTextMessage(ctx context.Context, b *bot.Bot, ChatID int64, text string, ReplyToMessageID ...int){
	messageParams := &bot.SendMessageParams{
		ChatID:      ChatID,
		Text:        text,
		ParseMode: models.ParseModeMarkdown,
	}

	if len(ReplyToMessageID) > 0 {
		messageParams.ReplyToMessageID = ReplyToMessageID[0]
	}

	_, err := b.SendMessage(ctx, messageParams)
	if err != nil {
		log.Panic(err)
	}
}

func SendCallbackQueryMessage(ctx context.Context, b *bot.Bot, queryID string, chatID int64, text string, _data ...any){

	callbackParams := &bot.AnswerCallbackQueryParams{
		CallbackQueryID: queryID,
		ShowAlert:       false,
	}


	messageParams := &bot.SendMessageParams{
		ChatID: chatID,
		Text: text,
	}

	if len(_data) > 0 {
		for _, value := range _data {
			switch value {
			case "callback_params":

			case "message_params":

			}
		}
	}

	b.AnswerCallbackQuery(ctx, callbackParams)

	_, err := b.SendMessage(ctx, messageParams)
	if err != nil {
		log.Panic(err)
	}

}