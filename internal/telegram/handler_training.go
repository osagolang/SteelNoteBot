package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleTraining(chatID int64) {

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноги", "legs"),
			tgbotapi.NewInlineKeyboardButtonData("Спина", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Грудные", "chest"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Плечи", "shoulders"),
			tgbotapi.NewInlineKeyboardButtonData("Бицепс", "biceps"),
			tgbotapi.NewInlineKeyboardButtonData("Трицепс", "triceps"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Икры", "calves"),
			tgbotapi.NewInlineKeyboardButtonData("Пресс", "press"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться в главное меню", "start"),
		),
	)

	h.HandleSendMessage(chatID, "Какую группу мышц будем тренировать? 👇", keyboard)

}
