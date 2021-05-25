package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command")

	switch message.Command() {
	case commandStart:
		msg.Text = "Start"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleGroupCommands(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command from group chat")
	switch message.Command() {
	case commandStart:
		msg.Text = "start from group chat"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel, groupOne, groupTwo, groupThree int64) {
	for update := range updates {

		if update.CallbackQuery != nil && update.CallbackQuery.Data == "asdasdasd" && update.CallbackQuery.Message.Chat.ID == groupOne {
			b.handleCallBackQueryHandler(update,  groupThree, groupOne)
		}

		if update.Message == nil { // ignore any non-Message Updates
			if update.ChannelPost == nil {
				continue
			} else {
				b.handleChannelPosts(update, groupOne, groupTwo)
			}
		} else if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
	}
}

func (b *Bot) handleCallBackQueryHandler(update tgbotapi.Update, groupThree, groupOne int64) {
	if update.CallbackQuery.From.UserName == "" {
		b.bot.Send(tgbotapi.NewMessage(groupOne, "⛔⛔⛔\n"+update.CallbackQuery.From.FirstName+" "+update.CallbackQuery.From.LastName+"\n"+"Пожалуйста создайте Username в телеграм акаунте!\n Затем нажмите снова кнопку!"))
	} else {
		userId, err := b.buttonRepo.GetUserId(update.CallbackQuery.Message.MessageID)
		if err != nil {
			log.Fatal(err)
		}

		if userId != update.CallbackQuery.From.ID {
			text, err := b.messageRepo.GetText(update.CallbackQuery.Message.MessageID)
			if err != nil {
				log.Fatal(err)
			}
			msg := tgbotapi.NewMessage(groupThree, "Принял: "+update.CallbackQuery.From.FirstName+" "+update.CallbackQuery.From.LastName+"\n"+"По заказу: "+text)

			button := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("Связаться", "https://t.me/"+update.CallbackQuery.From.UserName)))
			msg.ReplyMarkup = button
			err = b.buttonRepo.InsertButton(update.CallbackQuery.Message.MessageID, update.CallbackQuery.From.ID)
			if err != nil {
				log.Fatal(err)
			}

			b.bot.Send(msg)
		}
	}
}

func (b *Bot) handleChannelPosts(update tgbotapi.Update, groupOne, groupTwo int64) {
	if update.ChannelPost.Chat.ID == groupTwo {
		//if update.ChannelPost.Text[:12] == "Channel Post" {
		//	continue
		//}
		if len(update.ChannelPost.Text) > 70 {
			text := update.ChannelPost.Text[:23]
			if text == "Завод -Almatybeton" {
				arr := strings.Split(update.ChannelPost.Text, "\n")
				if len(arr) < 12 {
					msg := tgbotapi.NewMessage(groupTwo, "⛔Недостаточно информации⛔")
					msg.ReplyToMessageID = update.ChannelPost.MessageID
					b.bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(groupOne, arr[0]+"\n"+arr[3]+"\n"+arr[4]+"\n"+arr[9]+"\n"+arr[10]+"\n"+arr[12])
				msg.ReplyMarkup = numericKeyboard
				m, _ := b.bot.Send(msg)
				err := b.messageRepo.InsertMessage(m.MessageID, update.ChannelPost.Text)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				arr := strings.Split(update.ChannelPost.Text, "\n")
				if len(arr) < 21 {
					msg := tgbotapi.NewMessage(groupTwo, "⛔Недостаточно информации⛔")
					msg.ReplyToMessageID = update.ChannelPost.MessageID
					b.bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(groupOne, arr[3]+"\n"+arr[9]+"\n"+arr[20]+"\n"+arr[21]+"\n"+arr[4])
				msg.ReplyMarkup = numericKeyboard
				m, _ := b.bot.Send(msg)
				err := b.messageRepo.InsertMessage(m.MessageID, update.ChannelPost.Text)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			msg := tgbotapi.NewMessage(groupTwo, "⛔Недостаточно информации⛔")
			msg.ReplyToMessageID = update.ChannelPost.MessageID
			b.bot.Send(msg)
		}
	}
}
