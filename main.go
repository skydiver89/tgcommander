package main

import (
	"log"
	"os/exec"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mpvl/unique"
)

var config Config
var keyboard tgbotapi.ReplyKeyboardMarkup

func main() {
	err := config.load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	createKeyboard()

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
				if !userIsAllowed(update.Message.From.ID, bot, update.Message.Chat.ID) {
					continue
				}
				answer(update.Message.Text, bot, update.Message.Chat.ID)
			}
		}
	}
}

func createKeyboard() {
	var bufrows []int
	for _, button := range config.Buttons {
		bufrows = append(bufrows, button.Row)
	}
	unique.Ints(&bufrows)
	numrows := len(bufrows)
	for i := 0; i < numrows; i++ {
		var row []tgbotapi.KeyboardButton
		for _, button := range config.Buttons {
			if button.Row == i {
				row = append(row, tgbotapi.NewKeyboardButton(button.Name))
			}
		}
		keyboard.Keyboard = append(keyboard.Keyboard, row)
	}
}

func userIsAllowed(user int64, bot *tgbotapi.BotAPI, chat int64) bool {
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
	for _, button := range config.Buttons {
		if button.Name == message {
			processCommand(button, bot, chat)
			return
		}
	}
	sendKeyboard(bot, chat)
}

func processCommand(button Button, bot *tgbotapi.BotAPI, chat int64) {
	out, err := exec.Command(button.Command, button.Arguments...).Output()
	if err != nil {
		msg := tgbotapi.NewMessage(chat, err.Error())
		bot.Send(msg)
		return
	}
	if button.Output {
		msg := tgbotapi.NewMessage(chat, string(out))
		bot.Send(msg)
	}
}

func sendKeyboard(bot *tgbotapi.BotAPI, chat int64) {
	msg := tgbotapi.NewMessage(chat, "Choose action:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
