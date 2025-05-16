package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) HandleCallback(cb *tgbotapi.CallbackQuery) error {

	chatID := cb.Message.Chat.ID
	messageID := cb.Message.MessageID
	data := cb.Data

	h.HandleDeleteMessage(chatID, messageID)

	switch data {
	case "training":
		h.HandleTraining(chatID)
	case "legs", "back", "chest", "small":
		h.HandleExercise(chatID, data)
	default:
		h.HandleSendMessage(chatID, "Неизвестная команда", nil)
	}

	return nil
}
