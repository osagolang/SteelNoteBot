package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

func (h *Handler) HandleStart(msg *tgbotapi.Message) {

	chatID := msg.Chat.ID
	messageID := msg.MessageID

	h.HandleDeleteMessage(chatID, messageID)

	if lastMsgID, ok := h.lastMsgID[chatID]; ok {
		h.HandleDeleteMessage(chatID, lastMsgID)
		delete(h.lastMsgID, chatID)
	}

	user := &models.User{
		TelegramID: msg.From.ID,
		Username:   msg.From.UserName,
	}
	_ = h.userSVC.RegisterUser(context.Background(), user)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тренировка", "training"),
			tgbotapi.NewInlineKeyboardButtonData("Рекорды", "records"),
		),
	)

	mmm := h.HandleSendMessage(chatID, "Потренируемся или показать твои рекорды?\n\nВыбирай ниже 👇", keyboard)
	h.lastMsgID[chatID] = mmm
}
