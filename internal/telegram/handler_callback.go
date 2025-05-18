package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func (h *Handler) HandleCallback(cb *tgbotapi.CallbackQuery) error {

	chatID := cb.Message.Chat.ID
	messageID := cb.Message.MessageID
	data := cb.Data

	h.HandleDeleteMessage(chatID, messageID)

	switch {
	case data == "training":
		h.HandleTraining(chatID)
	case data == "legs" || data == "back" || data == "chest" || data == "small":
		h.HandleExercise(chatID, data)
	case strings.HasPrefix(data, "exercise_"):
		idString := strings.TrimPrefix(data, "exercise_")
		idExercise, err := strconv.Atoi(idString)
		if err != nil {
			h.HandleSendMessage(chatID, "Ошибка получения ID упражнения", nil)
			return nil
		}
		h.HandleExerciseSelected(chatID, idExercise)
	default:
		h.HandleSendMessage(chatID, "Неизвестная команда", nil)
	}

	return nil
}
