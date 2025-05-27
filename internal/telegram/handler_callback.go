package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"slices"
	"strconv"
	"strings"
)

func (h *Handler) HandleCallback(cb *tgbotapi.CallbackQuery) {

	chatID := cb.Message.Chat.ID
	messageID := cb.Message.MessageID
	data := cb.Data

	h.HandleDeleteMessage(chatID, messageID)

	muscleGroups := []string{"legs", "back", "chest", "shoulders", "biceps", "triceps", "calves"}

	switch {
	case data == "training":
		h.HandleTraining(chatID)
	case data == "records":
		h.HandleShowRecords(chatID)
	case slices.Contains(muscleGroups, data):
		h.HandleExerciseByGroup(chatID, data)
	case strings.HasPrefix(data, "exercise_"):
		idString := strings.TrimPrefix(data, "exercise_")
		idExercise, err := strconv.Atoi(idString)
		if err != nil {
			h.HandleSendMessage(chatID, "Ошибка получения ID упражнения", nil)
			return
		}
		h.HandleExerciseSelected(chatID, idExercise)
	default:
		h.HandleSendMessage(chatID, "Неизвестная команда", nil)
	}

	return
}
