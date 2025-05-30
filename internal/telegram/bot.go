package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/osagolang/SteelNoteBot/internal/config"
	"github.com/osagolang/SteelNoteBot/internal/services"
	"log"
)

// StartBot запускает телеграм бота
func StartBot(userSVC *services.UserService, exerciseSVC *services.ExerciseService, recordSVC *services.RecordService) {

	bot, err := tgbotapi.NewBotAPI(config.GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Авторизован как %s", bot.Self.UserName)

	handler := NewHandler(bot, userSVC, exerciseSVC, recordSVC)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обработка входящих сообщений
	for update := range updates {
		handler.HandleUpdate(update)
	}
}
