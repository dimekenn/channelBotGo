package telegram

import (
	"github.com/dimekenn/betonBot/pkg/telegram/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	messageRepo *repository.MessageRepo
	buttonRepo *repository.ButtonRepo
}

func NewBot(bot *tgbotapi.BotAPI, messageRepo *repository.MessageRepo,buttonRepo *repository.ButtonRepo) *Bot {
	return &Bot{bot: bot, messageRepo: messageRepo, buttonRepo: buttonRepo}
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Принимаю", "asdasdasd"),
	),
)

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	groupOne := int64(-1001311840611)
	groupTwo := int64(-1001401431774)
	groupThree := int64(-577413479)

	updates, err := b.bot.GetUpdatesChan(u)
	if err !=nil{
		log.Fatal(err)
	}

	b.handleUpdates(updates, groupOne, groupTwo, groupThree)
}


