package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleTraining(chatID int64) {

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноги", "legs"),
			tgbotapi.NewInlineKeyboardButtonData("Спина", "back"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Грудные", "chest"),
			tgbotapi.NewInlineKeyboardButtonData("Мелкие", "small"),
		),
	)

	h.HandleSendMessage(chatID, "Какую группу мышц будем тренировать? 👇", keyboard)

}
