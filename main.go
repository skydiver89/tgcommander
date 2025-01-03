package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mpvl/unique"
)

var VERSION = "0.1.0"
var GITREV = ""
var BUILDTIME = ""

var configFile = "config.yaml"
var config Config
var keyboard tgbotapi.ReplyKeyboardMarkup

func parseFlags() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-c configFile] [-h] [-v]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	var showHelp bool
	flag.BoolVar(&showHelp, "h", false, "Show this help")
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Show version information")
	flag.StringVar(&configFile, "c", "config.yaml", "Path to config file")
	flag.Parse()
	if showHelp {
		flag.Usage()
	}
	if showVersion {
		fmt.Println("Version      : ", VERSION)
		fmt.Println("Git revision : ", GITREV)
		fmt.Println("Build date   : ", BUILDTIME)
		os.Exit(0)
	}
}

func main() {
	parseFlags()
	err := config.load(configFile)
	if err != nil {
		log.Fatal(err)
	}
	createKeyboard()

	for {
		bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
		if err != nil {
			log.Println(err)
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
				go answer(update.Message, bot)
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
	msg := tgbotapi.NewMessage(chat, config.Telegram.DeclineMessage)
	bot.Send(msg)
	return false
}

func answer(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	defer bot.Request(tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID))
	for _, button := range config.Buttons {
		if button.Name == message.Text {
			processCommand(button, bot, message.Chat.ID)
			return
		}
	}
	sendKeyboard(bot, message.Chat.ID)
}

func processCommand(button Button, bot *tgbotapi.BotAPI, chat int64) {
	out, err := exec.Command(button.Command, button.Arguments...).Output()
	if err != nil {
		msg := tgbotapi.NewMessage(chat, err.Error())
		bot.Send(msg)
		return
	}
	if button.Output {
		res := fmt.Sprintf("Result of %s (%s %s):\n%s", button.Name, button.Command, strings.Join(button.Arguments, " "), string(out))
		msg := tgbotapi.NewMessage(chat, res)
		bot.Send(msg)
	}
}

func sendKeyboard(bot *tgbotapi.BotAPI, chat int64) {
	msg := tgbotapi.NewMessage(chat, "Choose action:")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
