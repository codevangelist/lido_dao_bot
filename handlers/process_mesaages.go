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

type ExtendedInlineKeyboardButton struct {
	models.InlineKeyboardButton
	ThumbURL string
}

func (e *ExtendedInlineKeyboardButton) AdditionalMethod() {
	
	// Add your custom logic here
}

var connectWallet = &models.InlineKeyboardMarkup{
	InlineKeyboard: [][]models.InlineKeyboardButton{
		{
			{Text: "Metamask", CallbackData: "metamask"},
			{Text: "WalletConnect", CallbackData: "wallet_connect"},
		}, 
		{
			{Text: "Coinbase", CallbackData: "coinbase"},
			// {Text: "Polkadot", CallbackData: "polkadot"},
		},
		{
			// {Text: "Kusama", CallbackData: "kusama"},
		},
	},
}



// var results = []models.InlineQueryResult{
// 	&models.InlineQueryResultArticle{ID: "1", Title: "Foo 1", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 1"}},
// 	&models.InlineQueryResultArticle{ID: "2", Title: "Foo 2", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 2"}},
// 	&models.InlineQueryResultArticle{ID: "3", Title: "Foo 3", InputMessageContent: &models.InputTextMessageContent{MessageText: "foo 3"}},
// }

func ProcessMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.InlineQuery == nil {
		return
	}

	results := []models.InlineQueryResult{
		&models.InlineQueryResultArticle{ID: "1", Title: "Become a Lido Node Operator", InputMessageContent: &models.InputTextMessageContent{
			MessageText: "foo 1",
			ParseMode: models.ParseModeMarkdown,
		}},
		&models.InlineQueryResultArticle{ID: "2", Title: "Stake with Lido", InputMessageContent: &models.InputTextMessageContent{
			MessageText: "foo 2",
			ParseMode: models.ParseModeMarkdown,
		}},
		&models.InlineQueryResultArticle{ID: "3", Title: "Claim Lido DAO tokens", InputMessageContent: &models.InputTextMessageContent{
			MessageText: "foo 3",
			ParseMode: models.ParseModeMarkdown,
		}},
	}

	b.AnswerInlineQuery(ctx, &bot.AnswerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		Results:       results,
	})

	switch update.Message.Text {
	case "/start":
		reply := "Welcome to"
		SendTextMessage(ctx, b, update.Message.Chat.ID, reply)

	case "/help":
		reply := "Welcome tocf"
		SendTextMessage(ctx, b, update.Message.Chat.ID, reply)

	case "/stake":
		reply := "Lido lets you stake tokens from many networks. Choose a network below to get started."
		SendInlineKeyboard(ctx, b, update.Message.Chat.ID, reply, supportedNetwors)


	case "/faq":
		
		// reply := "Lido lets you stake tokens from many networks. Choose a network below to get started."
		// SendInlineQuery(ctx, b, update.InlineQuery.ID, update.InlineQuery, results)
	
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
		
	case "coinbase":
		reply := "Callback Query reply to coin"
		SendCallbackQueryMarkUpMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply, connectWallet)

	case "connect_wallet":
		reply := "*Connect wallet*"
		SendCallbackQueryMarkUpMessage(ctx, b, update.CallbackQuery.ID, update.CallbackQuery.Message.Chat.ID, reply, connectWallet)
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

func SendCallbackQueryMessage(ctx context.Context, b *bot.Bot, queryID string, chatID int64, text string){

	callbackParams := &bot.AnswerCallbackQueryParams{
		CallbackQueryID: queryID,
		ShowAlert:       false,
	}


	messageParams := &bot.SendMessageParams{
		ChatID: chatID,
		Text: text,
	}

	b.AnswerCallbackQuery(ctx, callbackParams)

	_, err := b.SendMessage(ctx, messageParams)
	if err != nil {
		log.Panic(err)
	}
}

func SendCallbackQueryMarkUpMessage(ctx context.Context, b *bot.Bot, queryID string, chatID int64, text string, kb *models.InlineKeyboardMarkup){

	callbackParams := &bot.AnswerCallbackQueryParams{
		CallbackQueryID: queryID,
		ShowAlert:       false,
	}


	messageParams := &bot.SendMessageParams{
		ChatID: chatID,
		Text: text,
		ReplyMarkup: kb,
		ParseMode: models.ParseModeMarkdown,
	}

	b.AnswerCallbackQuery(ctx, callbackParams)

	_, err := b.SendMessage(ctx, messageParams)
	if err != nil {
		log.Panic(err)
	}
}

func SendInlineQuery(ctx context.Context, b *bot.Bot, queryID string, inlineQuery *models.InlineQuery, result []models.InlineQueryResult){
	if inlineQuery == nil {
		return
	}

	inlineQueryParams :=  &bot.AnswerInlineQueryParams{
		InlineQueryID: queryID,
		Results: result,
	}

	_, err := b.AnswerInlineQuery(ctx, inlineQueryParams)
	if err != nil{
		log.Panic(err)
	}
}