package bot

import (
	"gopkg.in/telebot.v3"
	"log"
	"time"

	"7-steelNote-bot/internal/config"
)

func StartBot() {

	token := config.GetToken()

	set := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 5 * time.Second},
	}

	bot, err := telebot.NewBot(set)
	if err != nil {
		log.Fatal(err)
	}

	buttonTrain := telebot.Btn{Text: "Тренировка"}
	buttonRecord := telebot.Btn{Text: "Рекорды"}

	menu := &telebot.ReplyMarkup{}
	menu.Reply(
		menu.Row(buttonTrain, buttonRecord),
	)

	bot.Handle("/start", func(c telebot.Context) error {
		_ = c.Delete()
		return c.Send("Хочешь посмотреть свои рекорды или записать тренировку? Выбирай ниже 👇", menu)
	})

	bot.Handle(&buttonTrain, func(c telebot.Context) error {
		_ = c.Delete()
		return c.Send("Отлично, сейчас будем тренироваться. (дальше тут будет выбор группы мышц для тренировки...)")
	})

	bot.Handle(&buttonRecord, func(c telebot.Context) error {
		return c.Send("Тут буду отображены твои рекорды и достижения...")
	})

	log.Println("Bot is running")
	bot.Start()
}
