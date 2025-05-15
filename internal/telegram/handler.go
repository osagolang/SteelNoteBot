package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/osagolang/SteelNoteBot/internal/services"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	userSVC *services.UserService
}

func NewHandler(bot *tgbotapi.BotAPI, userSVC *services.UserService) *Handler {
	return &Handler{bot: bot, userSVC: userSVC}
}

func (h *Handler) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		switch update.Message.Text {
		case "/start":
			h.HandleStart(update.Message)
		}
	}
}
