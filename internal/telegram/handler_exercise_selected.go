package telegram

import (
	"context"
	"fmt"
)

func (h *Handler) HandleExerciseSelected(chatID int64, exerciseID int) {

	records, err := h.recordSVC.GetRecords(context.Background(), chatID, exerciseID, 5)
	if err != nil {
		h.HandleSendMessage(chatID, "Ошибка получения тренировок", nil)
		return
	}

	msg := ""

	if len(records) == 0 {
		msg = "Упс! Ты ещё ничего не записывал\n"
	} else {
		msg = "Последние тренировки:\n\n"

		for _, r := range records {
			msg += fmt.Sprintf("%s - %.1f кг. x %d раз(а)\n", r.CreatedAt.Format("02.01"), r.Weight, r.Reps)
		}
	}

	msg += "\n---\nВведи данные текущей тренировки\nв формате: вес повторения (например: 90 10)"

	h.tempInput[chatID] = exerciseID

	h.HandleSendMessage(chatID, msg, nil)

}
