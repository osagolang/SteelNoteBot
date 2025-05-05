package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN not found in .env")
	}

	param := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 5 * time.Second},
	}

	bot, err := telebot.NewBot(param)
	if err != nil {
		log.Fatal(err)
	}

	battonTrain := telebot.Btn{Text: "Тренировка"}
	battonRecord := telebot.Btn{Text: "Рекорды"}

	menu := &telebot.ReplyMarkup{}
	menu.Reply(
		menu.Row(battonTrain, battonRecord),
	)

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Привет! Я Бот - Железный Блокнот. Жмякай на кнопки ниже!", menu)
	})

	bot.Handle(&battonTrain, func(c telebot.Context) error {
		return c.Send("Отлично, сейчас будем тренироваться. (дальше тут будет выбор группы мышц для тренировки...)")
	})

	bot.Handle(&battonRecord, func(c telebot.Context) error {
		return c.Send("Тут буду отображены твои рекорды и достижения...")
	})

	log.Println("Bot is running")
	bot.Start()

}
