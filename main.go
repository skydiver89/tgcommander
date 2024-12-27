package main

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var config Config

func main() {
	err := config.load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for {
		bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
		if err != nil {
			log.Fatal(err)
			time.Sleep(time.Second * 10)
			continue
		}
		bot.Debug = false
		log.Println("Authorised on account", bot.Self.UserName)
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)
		for update := range updates {
			if update.Message != nil {
				fmt.Println("Message from ", update.Message.From.UserName)
				if !userIsAllowed(update.Message.From.UserName, bot, update.Message.Chat.ID) {
					continue
				}
				answer(update.Message.Text, bot, update.Message.Chat.ID)
			}
		}
	}
}

func userIsAllowed(user string, bot *tgbotapi.BotAPI, chat int64) bool {
	for _, u := range config.Telegram.Users {
		if u == user {
			return true
		}
	}
	msg := tgbotapi.NewMessage(chat, "Ты кто такой? Давай, до свидания!")
	bot.Send(msg)
	return false
}

func answer(message string, bot *tgbotapi.BotAPI, chat int64) {
}
