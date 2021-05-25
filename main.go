package main

import (
	"database/sql"
	"github.com/dimekenn/betonBot/pkg/telegram"
	"github.com/dimekenn/betonBot/pkg/telegram/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	"log"
)

func main()  {
	bot, err := tgbotapi.NewBotAPI("1644632728:AAFoYix8mgvmcL689DLpITDVTRyF3v18Ayw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	db, err := dbInit("postgres", "postgres://postgres:mypassword@localhost:5432/postgres?sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	messageRepo := repository.NewMessageRepo(db)
	buttonRepo := repository.NewButtonRepo(db)

	telegramBot := telegram.NewBot(bot, messageRepo, buttonRepo)

	telegramBot.Start()
}

func dbInit(name, url string) (*sql.DB, error) {
	db, err := sql.Open(name, url)
	if err !=nil{
		return nil, err
	}
	return db, nil
}